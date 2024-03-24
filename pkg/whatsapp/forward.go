package whatsapp

import (
	"context"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type FormatterInterface interface {
	FormatString(jid types.JID, conv string) *string
	DecodeMessage(conv string) (jid types.JID, res string, err error)
}

type Message struct {
	client    *whatsmeow.Client
	dest      types.JID
	formatter FormatterInterface
}

func NewMessage(client *whatsmeow.Client, dest types.JID, formatter FormatterInterface) *Message {
	return &Message{client, dest, formatter}
}

func (m *Message) FormatText(ctx context.Context, jid types.JID, conv string) {
	m.Send(ctx, &waProto.Message{Conversation: m.formatter.FormatString(jid, conv)})
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
