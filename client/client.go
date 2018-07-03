package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")

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
