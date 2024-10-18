package server

import (
	"github.com/google/uuid"
	"net"
)

type Client struct {
	channel chan ProcessorData
	id      string
	roomId  string
	conn    net.Conn
	room    *Room
	ronin   bool
}

func InitClient(room string, conn net.Conn) *Client {
	return &Client{
		id:     uuid.New().String(),
		roomId: room,
		conn:   conn,
		ronin:  true,
	}
}

type Room struct {
	channel chan ProcessorData
	Id      string
	clients []*Client
}

func InitRoom(channel chan ProcessorData, clients []*Client) *Room {
	return &Room{
		Id:      uuid.New().String(),
		channel: channel,
		clients: clients,
	}
}

type ProcessorData struct {
	client  *Client
	rawData string
}
