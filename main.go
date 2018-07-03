package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var clientNodes map[*Client]net.Addr

type Client struct {
	address    net.Addr
	outbound   chan string
	conn       net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
	connection *Client
}

func (client *Client) Read() {
	for {
		message, err := client.reader.ReadString('\n')
		fmt.Println(len(clientNodes))
		if err != nil {
			return
		}
		fmt.Print("From client "+client.address.String()+": ", string(message))
		for node := range clientNodes {
			if node.address != client.address {
				node.outbound <- message
			}
		}
	}
}

func (client *Client) Write() {
	for data := range client.outbound {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func CreateClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	client := &Client{
		outbound: make(chan string),
		conn:     conn,
		reader:   reader,
		writer:   writer,
		address:  conn.RemoteAddr(),
	}

	client.Listen()

	return client
}

func main() {
	clientNodes = make(map[*Client]net.Addr)
	fmt.Println("Server starting...")
	ln, _ := net.Listen("tcp", ":8081")
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			os.Exit(1)
		}
		client := CreateClient(conn)
		clientNodes[client] = client.address
		fmt.Println("New connection from: " + client.address.String())
	}
}
