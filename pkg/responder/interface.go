package responder

import (
	"context"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type MessageInterface interface {
	FormatText(ctx context.Context, jid types.JID, conv string)
	Dest() types.JID
	IsDest(jid types.JID) bool
	Send(ctx context.Context, msg *waProto.Message)
}

type Regex interface {
	FindString(string) string
	MatchString(string) bool
}

type FormatterInterface interface {
	FormatString(jid types.JID, conv string) *string
	DecodeMessage(conv string) (jid types.JID, res string, err error)
}
