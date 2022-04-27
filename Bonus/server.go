package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":11037");
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		for {
			r := bufio.NewReader(connection)

			for {
				in, err := r.ReadBytes(byte('\n'))
				if err != nil {
					break
				}

				in = append([]byte("Hello "), in...)

				connection.Write(in)
			}
		}
	}
}