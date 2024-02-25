package message

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"regexp"
	"responder/pkg/responder/formatter"
	"responder/pkg/responder/forwarder"
	"responder/pkg/responder/helper"
)

func EventListener(client *whatsmeow.Client, jidParser, convParser *regexp.Regexp, forwarder *forwarder.Message) func(event any) {
	return func(event any) {
		switch e := event.(type) {
		case *events.Message:
			EventConsumer(client, jidParser, convParser, forwarder, e)
		}
	}
}

func EventConsumer(client *whatsmeow.Client, jidParser, convParser *regexp.Regexp, forwarder *forwarder.Message, event *events.Message) {
	if event.Info.IsFromMe {
		return
	}
	var jid types.JID
	var err error
	defer func() {
		if err != nil {
			client.Log.Warnf("Sending Message Error to %s: %+v", jid, err)
		}
	}()

	if !forwarder.IsDest(event.Info.Chat) {
		forwarder.FormatText(context.Background(), event.Info.Chat, helper.GetMessage(event.Message))
		return
	}

	jid, conv, err := formatter.Decoder(jidParser, convParser, helper.GetMessage(event.Message))
	fmt.Printf("JID: %+v\n", jid)
	if err != nil {
		forwarder.FormatText(context.Background(), event.Info.Chat, err.Error())
		err = nil
		return
	}

	_, err = client.SendMessage(context.Background(), jid, &waProto.Message{Conversation: &conv})
}
