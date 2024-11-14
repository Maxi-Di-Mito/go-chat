package client

import "strings"

const MESSAGE_JOIN_ROOM = "MESSAGE-joinRoom"

func CreateMessageJoinRoom(roomId string) string {
	msg := []string{MESSAGE_JOIN_ROOM}
	msg = append(msg, roomId)

	return strings.Join(msg, "\n")
}

const MESSAGE_SEND_MESSAGE = "MESSAGE-sendMessage"

func CreateMessageSendMessage(text string) string {
	msg := []string{MESSAGE_JOIN_ROOM}
	msg = append(msg, text)

	return strings.Join(msg, "\n")
}
