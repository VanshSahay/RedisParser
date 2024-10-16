package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"redis-go/cmd"
	"redis-go/redisparser"
)

func main() {
	fmt.Println("Logs from your program will appear here!")
	
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	for{
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		
		go handleConn(c)
	}
}

func handleConn(c net.Conn){
	defer c.Close()	
	buf := make([]byte, 128)

	for {
		_, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return
			}
			fmt.Println("Error reading buffer: ", err.Error())
			os.Exit(1)
		}
	
		log.Printf("read command:\n%s", buf)

		command, err := redisparser.ParseObject(buf)
		if err!=nil {
			fmt.Println("Error parsing command: ", err.Error())
		}

		// fmt.Printf("RAW COMMAND: %q\n", command)		
		switch command[1] {
		case "ECHO":
			c.Write([]byte(command[2]+"\r\n"))
		case "SET":
			status := cmd.SetCMD(command[2], command[3])
			c.Write([]byte(status + "\r\n"))
		case "GET":
			value, err := cmd.GetCMD(command[2])
			if len(value)>0 {
				c.Write([]byte(value + "\r\n"))
			} else {
				fmt.Println("Error getting values: ", err)
			}
		default:
			fmt.Println("INVALID COMMAND")
		}
	}
}