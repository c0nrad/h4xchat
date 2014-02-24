package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

type Message struct {
	text string
	from net.Conn
}

func main() {
	var port string
	flag.StringVar(&port, "port", "1337", "Port to host server on")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	msgChan := make(chan Message)
	addChan := make(chan Client)
	rmChan := make(chan Client)

	go handleMessages(msgChan, addChan, rmChan)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn, msgChan, addChan, rmChan)
	}
}

func handleConnection(c net.Conn, msgChan chan<- Message, addChan chan<- Client, rmChan chan<- Client) {
	addChan <- Client{c}

	buf := bufio.NewReader(c)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		if string(line) == "" {
			continue
		}
		msgChan <- Message{string(line), c}
	}

	c.Close()
	fmt.Printf("Connection closed")
	rmChan <- Client{c}
}

func handleMessages(msgChan <-chan Message, addChan <-chan Client, rmChan <-chan Client) {
	clients := make(map[net.Conn]Client)
	for {
		select {
		case message := <-msgChan:
			fmt.Println("Recieved message", message)
			broadcastMessage(clients, message)
		case client := <-addChan:
			fmt.Printf("New client: %v\n", client.conn)
			clients[client.conn] = client
		case client := <-rmChan:
			fmt.Printf("Client disconnects: %v\n", client.conn)
			delete(clients, client.conn)
		}
	}
}

func broadcastMessage(clients map[net.Conn]Client, msg Message) {
	for conn, _ := range clients {
		if conn != msg.from {
			go func(conn net.Conn, msg Message) {
				fmt.Fprintf(conn, msg.text+"\n")
			}(conn, msg)
		}
	}
}
