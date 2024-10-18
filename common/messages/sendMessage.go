package messages

import (
	"strings"
)

const MESSAGE_DELIVER_MESSAGE = "MESSAGE-deliverMessage"

func CreateMessageDeliverMessage(text string, fromId string) string {
	msg := []string{MESSAGE_DELIVER_MESSAGE}

	msg = append(msg, "FROM:"+fromId)
	msg = append(msg, text)

	return strings.Join(msg, "\n")
}
