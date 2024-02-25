package helper

import (
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func GetMessage(Message *waProto.Message) (message string) {
	switch {
	case Message.ExtendedTextMessage != nil:
		return Message.ExtendedTextMessage.GetText()
	case Message.DocumentMessage != nil:
		return Message.DocumentMessage.GetCaption()
	case Message.ImageMessage != nil:
		return Message.ImageMessage.GetCaption()
	case Message.LiveLocationMessage != nil:
		return Message.LiveLocationMessage.GetCaption()
	default:
		return Message.GetConversation()
	}
}
