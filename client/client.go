package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println()
			if err != nil {
				fmt.Println("TCP connection closed.")
				os.Exit(0)
			}
			fmt.Print("Receiving from server: " + message)
			fmt.Println(conn)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Sending to server: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
