package server

import (
	"github.com/google/uuid"
)

type Client struct {
	channel chan string
	id      string
	roomId  string
}

func (c *Client) InitClient(room string, channel chan string) *Client {
	return &Client{
		id:      uuid.New().String(),
		roomId:  room,
		channel: channel,
	}
}

type Room struct {
	channel   chan string
	id        string
	clientOne *Client
	clientTwo *Client
}
