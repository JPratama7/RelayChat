package formatter

import (
	"errors"
	"fmt"
	"github.com/JPratama7/util/sync"
	"go.mau.fi/whatsmeow/types"
	"responder/pkg/responder"
	"strings"
)

type Format struct {
	pool                       *sync.Pool[*strings.Builder]
	jidC, convC                responder.Regex
	headerFormat, messageSplit string
}

func NewFormat(headerFormat, msgSplit string, jid, conv responder.Regex) *Format {
	return &Format{
		pool: sync.NewPool(func() *strings.Builder {
			return new(strings.Builder)
		}),
		headerFormat: headerFormat,
		messageSplit: msgSplit,
		jidC:         jid,
		convC:        conv,
	}
}

func (r *Format) FormatString(jid types.JID, conv string) *string {
	builder := r.pool.Get()
	defer func() {
		r.pool.Put(builder)
		builder.Reset()
	}()

	builder.WriteString(fmt.Sprintf("%s:%s\n", r.headerFormat, jid.User))
	builder.WriteString(r.messageSplit)
	builder.WriteRune('\n')
	builder.WriteString(conv)
	result := builder.String()
	return &result
}

func (r *Format) DecodeMessage(conv string) (jid types.JID, res string, err error) {
	if !r.jidC.MatchString(conv) {
		err = errors.New("JID not found")
		return
	}

	if !r.convC.MatchString(conv) {
		err = errors.New("conversation not found")
		return
	}

	jid = types.NewJID(strings.ReplaceAll(strings.ReplaceAll(r.jidC.FindString(conv), " ", ""), r.headerFormat, ""), types.DefaultUserServer)

	return
}

func Decoder(jidC, convC responder.Regex, conv string) (jid types.JID, res string, err error) {

	if !jidC.MatchString(conv) {
		err = errors.New("JID not found")
		return
	}

	if !convC.MatchString(conv) {
		err = errors.New("conversation not found")
		return
	}

	jid = types.NewJID(strings.ReplaceAll(strings.ReplaceAll(jidC.FindString(conv), " ", ""), "Sender:", ""), types.DefaultUserServer)
	resString := strings.Split(convC.FindString(conv), "----Message----\n")
	res = resString[len(resString)-1]
	return

}
