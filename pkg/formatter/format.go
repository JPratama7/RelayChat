package formatter

import (
	"errors"
	"fmt"
	"github.com/JPratama7/util/sync"
	"go.mau.fi/whatsmeow/types"
	"strings"
)

type Regex interface {
	FindStringSubmatch(string) []string
	MatchString(string) bool
}

type Format struct {
	pool                       *sync.Pool[*strings.Builder]
	jidC, convC                Regex
	headerFormat, messageSplit string
}

func NewFormat(headerFormat, msgSplit string, jid, conv Regex) *Format {
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
	fmt.Println(conv)

	jidString := r.jidC.FindStringSubmatch(conv)
	if len(jidString) < 1 {
		err = errors.New("regex doesnt match")
		return
	}
	stringMatch := r.convC.FindStringSubmatch(conv)
	if len(stringMatch) < 1 {
		err = errors.New("regex doesnt match")
		return
	}

	jid = types.NewJID(jidString[1], types.DefaultUserServer)
	res = stringMatch[1]
	return
}

func Decoder(jidC, convC Regex, conv string) (jid types.JID, res string, err error) {

	if !jidC.MatchString(conv) {
		err = errors.New("JID not found")
		return
	}

	if !convC.MatchString(conv) {
		err = errors.New("conversation not found")
		return
	}

	jid = types.NewJID(jidC.FindStringSubmatch(conv)[1], types.DefaultUserServer)
	resString := strings.Split(jidC.FindStringSubmatch(conv)[1], "----Message----\n")
	res = resString[len(resString)-1]
	return

}
