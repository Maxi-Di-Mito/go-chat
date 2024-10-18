package server

import (
	"errors"
	"fmt"
)

func JoinClientToRoom(roomId string, client *Client) {
	var roomObject *Room
	fmt.Println(roomId)
	if roomId == "0" {
		fmt.Println("creating new room")
		channel := make(chan ProcessorData, 10)
		roomObject = InitRoom(channel, []*Client{client})

	} else {
		fmt.Println("searching room")
		for idx, room := range Rooms {
			if room.Id == roomId {
				roomObject = &Rooms[idx]
				break
			}
		}
	}

	if roomObject == nil {
		panic(errors.New("no roomId" + roomId))
	}

	for idx, c := range roninClients {
		if c.id == client.id {
			client.channel = roomObject.channel
			client.ronin = false
			roninClients = append(roninClients[:idx], roninClients[idx+1:]...)
			roomObject.clients = append(roomObject.clients, &roninClients[idx])
			return
		}
	}
	panic(errors.New("no client" + client.id))
}

func deliverNewMessage(text string, from *Client) {
	var room *Room
	for idx, r := range Rooms {
		if r.Id == from.roomId {
			room = &Rooms[idx]
			break
		}
	}
	if room == nil {
		panic(errors.New("NO ROOM WITH ID" + from.roomId))
	}

	for _, client := range room.clients {
		if client.id != from.id {
			client.conn.Write([]byte(text))
		}
	}
}
