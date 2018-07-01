package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", os.Args[1]+":8081")
	if err != nil {
		os.Exit(1)
	}

	go func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Receiving from server: " + message)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Sending to server: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
