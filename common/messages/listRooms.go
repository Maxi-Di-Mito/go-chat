package messages

import (
	"strings"
)

const MESSAGE_LIST_ROOMS = "MESSAGE-listRooms"

func CreateMessageListRooms(roomIds []string) string {
	messages := []string{MESSAGE_LIST_ROOMS}
	messages = append(messages, "These are the rooms available")
	messages = append(messages, "new room: 0")
	for _, room := range roomIds {
		messages = append(messages, "Room ID: ", room)
	}

	messages = append(messages, "SEND THE ROOM ID YOU WANT TO JOIN TO")

	message := strings.Join(messages, "\n")
	return message
}
