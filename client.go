package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "172.16.45.76:8081")
	if err != nil {
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Sending to server: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Receiving from server: " + message)
	}
}
