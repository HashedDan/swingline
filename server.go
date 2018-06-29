package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Server starting...")
	ln, _ := net.Listen("tcp", ":8081")

	conn, err := ln.Accept()
	if err != nil {
		os.Exit(1)
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("From client: ", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
