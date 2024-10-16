# Redis Parser in Go

A fast, efficient tool for parsing Redis protocol messages. This lightweight Go implementation enables quick parsing of Redis commands and responses, supporting both synchronous and asynchronous operations for versatile integration with Redis-based applications.

## Features

-   Parses Redis protocol messages quickly and efficiently
-   Supports both synchronous and asynchronous parsing
-   Lightweight and easy to integrate
-   Written in Go for high performance
-   Implements basic Redis commands: ECHO, SET, GET

## Installation

To install the Redis Parser, use the following command:

```
go get github.com/VanshSahay/RedisParser
```

## Usage

Here's how to use the Redis Parser in your Go application:

1. Import the necessary packages:

```go
import (
    "github.com/VanshSahay/RedisParser"
)
```

2. Set up a TCP server:

```go
l, err := net.Listen("tcp", "0.0.0.0:6379")
if err != nil {
    fmt.Println("Failed to bind to port 6379")
    os.Exit(1)
}
defer l.Close()
```

3. Accept connections and handle them:

```go
for {
    c, err := l.Accept()
    if err != nil {
        fmt.Println("Error accepting connection: ", err.Error())
        continue
    }
    go handleConn(c)
}
```

4. In the `handleConn` function, read from the connection, parse the command, and execute it:

```go
func handleConn(c net.Conn) {
    defer c.Close()
    buf := make([]byte, 128)

    for {
        _, err := c.Read(buf)
        if err != nil {
            // Handle error
        }

        command, err := redisparser.ParseObject(buf)
        if err != nil {
            fmt.Println("Error parsing command: ", err.Error())
        }

        switch command[1] {
        case "ECHO":
            c.Write([]byte(command[2] + "\r\n"))
        case "SET":
            status := cmd.SetCMD(command[2], command[3])
            c.Write([]byte(status + "\r\n"))
        case "GET":
            value, err := cmd.GetCMD(command[2])
            if len(value) > 0 {
                c.Write([]byte(value + "\r\n"))
            } else {
                fmt.Println("Error getting values: ", err)
            }
        default:
            fmt.Println("INVALID COMMAND")
        }
    }
}
```

This setup creates a basic Redis server that can handle ECHO, SET, and GET commands.
