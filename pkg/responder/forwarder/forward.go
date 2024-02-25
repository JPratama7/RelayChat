package forwarder

import (
	"context"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type Message struct {
	client    *whatsmeow.Client
	dest      types.JID
	formatter func(types.JID, string) *string
}

func NewMessage(client *whatsmeow.Client, dest types.JID, formatter func(types.JID, string) *string) *Message {
	return &Message{client, dest, formatter}
}

func (m *Message) FormatText(ctx context.Context, jid types.JID, conv string) {
	m.Send(ctx, &waProto.Message{Conversation: m.formatter(jid, conv)})
}

func (m *Message) Dest() types.JID {
	return m.dest
}

func (m *Message) IsDest(jid types.JID) bool {
	return m.dest.User == jid.User && m.dest.Server == jid.Server

}

func (m *Message) Send(ctx context.Context, msg *waProto.Message) {
	if _, err := m.client.SendMessage(ctx, m.dest, msg); err != nil {
		m.client.Log.Warnf("Sending Message Error: %+v", err)
	}
}
