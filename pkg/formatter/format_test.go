package formatter

import (
	"github.com/stretchr/testify/assert"
	"go.mau.fi/whatsmeow/types"
	"regexp"
	"testing"
)

func TestFormatString(t *testing.T) {
	format := NewFormat("Sender", "----Message----", regexp.MustCompile("^Sender: (.*)"), regexp.MustCompile("----Message----\n(.*)"))
	jid := types.NewJID("testuser", types.DefaultUserServer)
	message := "Hello, World!"

	result := format.FormatString(jid, message)

	expected := "Sender:testuser\n----Message----\nHello, World!"
	assert.Equal(t, expected, *result)
}

func TestDecodeMessageWithValidInput(t *testing.T) {
	format := NewFormat("Sender", "----Message----", regexp.MustCompile("^Sender: (.*)"), regexp.MustCompile("----Message----\n(.*)"))
	message := "Sender: testuser\n----Message----\nHello, World!"

	jid, res, err := format.DecodeMessage(message)

	assert.NoError(t, err)
	assert.Equal(t, "testuser", jid.User)
	assert.Equal(t, "Hello, World!", res)
}

func TestDecodeMessageWithInvalidJID(t *testing.T) {
	format := NewFormat("Sender", "----Message----", regexp.MustCompile("^Sender: (.*)"), regexp.MustCompile("----Message----\n(.*)"))
	message := "Invalid: testuser\n----Message----\nHello, World!"

	_, _, err := format.DecodeMessage(message)

	assert.Error(t, err)
	assert.Equal(t, "JID not found", err.Error())
}

func TestDecodeMessageWithInvalidConversation(t *testing.T) {
	format := NewFormat("Sender", "----Message----", regexp.MustCompile("^Sender: (.*)"), regexp.MustCompile("----Message----\n(.*)"))
	message := "Sender: testuser\nInvalid\nHello, World!"

	_, _, err := format.DecodeMessage(message)

	assert.Error(t, err)
	assert.Equal(t, "conversation not found", err.Error())
}
