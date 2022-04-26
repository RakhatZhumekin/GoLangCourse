package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":11037");
	if err != nil {
		fmt.Println(err.Error())
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
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