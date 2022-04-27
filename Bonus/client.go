package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", ":11037")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(connection)

	for {
		fmt.Print("Write your name: ")
		userLine, err := input.ReadBytes(byte('\n'))
		if err != nil {
			break
		}

		connection.Write(userLine)

		serverLine, err := response.ReadBytes(byte('\n'))
		if err != nil {
			break
		}

		fmt.Println("Response: " + string(serverLine))
	}
}