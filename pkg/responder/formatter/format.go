package formatter

import (
	"errors"
	"github.com/JPratama7/util/sync"
	"go.mau.fi/whatsmeow/types"
	"regexp"
	"strings"
)

type Format struct {
	pool *sync.Pool[*strings.Builder]
}

func NewFormat() *Format {
	return &Format{pool: sync.NewPool(func() *strings.Builder {
		return new(strings.Builder)
	})}
}

func (r Format) FormatString(jid types.JID, conv string) *string {
	builder := r.pool.Get()
	defer func() {
		r.pool.Put(builder)
		builder.Reset()
	}()

	builder.WriteString("Sender: " + jid.User + "\n")
	builder.WriteString("----Message----")
	builder.WriteRune('\n')
	builder.WriteString(conv)
	result := builder.String()
	return &result

}
func Decoder(jidC, convC *regexp.Regexp, conv string) (jid types.JID, res string, err error) {

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
