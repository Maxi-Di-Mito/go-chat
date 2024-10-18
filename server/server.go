package server

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/Maxi-Di-Mito/go-routines/client"
	"github.com/Maxi-Di-Mito/go-routines/common/messages"
	"github.com/Maxi-Di-Mito/go-routines/utils"
)

var roninClients []Client

var Rooms []Room

var group sync.WaitGroup

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	connectionListener(listener)
	group.Wait()
}

func connectionListener(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR LISTENER", err)
			continue
		}

		client := NewRoningClient(conn)
		group.Add(1)
		go receiver(client)

		mess := messages.CreateMessageListRooms(utils.Map(Rooms, func(r Room) string {
			return r.Id
		}))

		conn.Write([]byte(mess))
	}
}

func NewRoningClient(conn net.Conn) *Client {
	newClient := InitClient("noRoom", conn)
	roninClients = append(roninClients, *newClient)
	return newClient
}

func receiver(client *Client) {
	defer client.conn.Close()

	in := make([]byte, 1024)
	for {
		_, err := client.conn.Read(in)
		if err != nil {
			fmt.Println("ERROR", err)
			break
		}
		msg := strings.TrimSuffix(string(in), "\n")

		fmt.Println("FROM: ", client.id, "MESSAGE", msg)

		if client.ronin {
			JoinClientToRoom(msg, client)
		} else {
			parseMessage(client, string(in))
		}

	}
	group.Done()
}

func parseMessage(from *Client, msg string) {
	msgLines := strings.Split(msg, "\n")

	header := msgLines[0]

	switch header {
	case client.MESSAGE_SEND_MESSAGE:
		deliverNewMessage(strings.Join(msgLines[1:], ""), from)

	default:
		fmt.Println("\n=============unknown message\n", msg, "\n=============")
	}
}
