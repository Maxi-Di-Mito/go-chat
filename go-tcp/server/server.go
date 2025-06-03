package server

import (
	"fmt"
	"net"
	"strings"

	"github.com/google/uuid"
)

var clientList = []*Client{}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	go broadCaster()
	listenerLoop(listener)
}

func listenerLoop(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR LISTENER", err)
			continue
		}
		client := &Client{
			conn: conn,
			id:   uuid.NewString(),
		}

		go clientLoop(client)
	}
}

var broadCastChannel = make(chan string, 100)

func clientLoop(client *Client) {
	defer client.conn.Close()
	clientList = append(clientList, client)
	in := make([]byte, 1024)
	for {
		n, err := client.conn.Read(in)
		if err != nil {
			fmt.Println("ERROR WITH CLIENT", err)
			break
		}
		msg := strings.TrimSpace(string(in[:n]))
		fmt.Println("MSG", msg)
		if client.name == "" {
			client.name = msg
			broadCastChannel <- fmt.Sprintf("%s has joined the chat", client.name)
		} else {
			broadCastChannel <- fmt.Sprintf("%s | %s :\n%s", client.name, client.id, msg)
		}
	}
}

func broadCaster() {
	for {
		msg := <-broadCastChannel
		for _, client := range clientList {
			client.conn.Write([]byte(msg))
		}
	}
}
