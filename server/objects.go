package server

import (
	"fmt"
	"github.com/google/uuid"
	"net"
)

// ///////////////////////////////////////////////////////////////
// Client Object and logic
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

// ///////////////////////////////////////////////////////////////
// Room Object and logic
type Room struct {
	channel chan ProcessorData
	Id      string
	clients []*Client
}

func (r *Room) Broadcast(msg string) {
	for _, c := range r.clients {
		c.conn.Write([]byte(msg))
	}
}

func (r *Room) QueueMsg(client *Client, msg string) {
	r.channel <- ProcessorData{
		client:  client,
		rawData: msg,
	}
}

func InitRoom(channel chan ProcessorData, clients []*Client) *Room {
	room := &Room{
		Id:      uuid.New().String(),
		channel: channel,
		clients: clients,
	}

	go RoomProcesor(room)
	return room
}

func RoomProcesor(room *Room) {
	for {
		in := <-room.channel
		room.Broadcast(fmt.Sprintf("%s :\n%s", in.client.id, in.rawData))
	}
}

type ProcessorData struct {
	client  *Client
	rawData string
}
