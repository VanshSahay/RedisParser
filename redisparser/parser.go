package redisparser

import (
	"fmt"
	"strings"
)

func ParseObject(buf []byte) ([]string, error) {
	errorArg := make([]string, 1)
	symbol := string(buf)[0]
	switch symbol {
	case '+':
		return parseSimpleString(buf)
	case '*':
		return parseArray(buf)
	default:
		return errorArg, fmt.Errorf("unknown type: %c", symbol)
	}
}

func parseSimpleString(buf []byte) ([]string, error) {
	bufCommand := string(buf)[1:]
	cleanCommand := strings.ReplaceAll(bufCommand, "\x00", "")
	command := strings.TrimRight(cleanCommand, "\r\n")
	commandArgs := []string{command}

	return commandArgs, nil
}

func parseArray(buf []byte) ([]string, error) {
	var commandArgs []string
	if len(buf) != 0 {
		bufCommand := string(buf)[1:]
		cleanCommand := strings.ReplaceAll(bufCommand, "\x00", "")
		parts := strings.Split(cleanCommand, "\r\n")
	
	
		for _, str := range parts {
			if str != "" {
				commandArgs = append(commandArgs, str)
			}
		}
	}else{
		return []string{""}, nil
	}

	return commandArgs, nil
}