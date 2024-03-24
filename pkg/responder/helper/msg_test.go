package helper

import (
	"github.com/JPratama7/util/convert"
	"github.com/stretchr/testify/assert"
	"go.mau.fi/whatsmeow/binary/proto"
	"testing"
)

func TestGetMessageFromExtendedTextMessage(t *testing.T) {
	message := &proto.Message{
		ExtendedTextMessage: &proto.ExtendedTextMessage{
			Text: convert.ToPointer("Hello, World!"),
		},
	}

	result := GetMessage(message)

	assert.Equal(t, "Hello, World!", result)
}

func TestGetMessageFromDocumentMessage(t *testing.T) {
	message := &proto.Message{
		DocumentMessage: &proto.DocumentMessage{
			Caption: convert.ToPointer("Document Caption"),
		},
	}

	result := GetMessage(message)

	assert.Equal(t, "Document Caption", result)
}

func TestGetMessageFromImageMessage(t *testing.T) {
	message := &proto.Message{
		ImageMessage: &proto.ImageMessage{
			Caption: convert.ToPointer("Image Caption"),
		},
	}

	result := GetMessage(message)

	assert.Equal(t, "Image Caption", result)
}

func TestGetMessageFromLiveLocationMessage(t *testing.T) {
	message := &proto.Message{
		LiveLocationMessage: &proto.LiveLocationMessage{
			Caption: convert.ToPointer("Live Location Caption"),
		},
	}

	result := GetMessage(message)

	assert.Equal(t, "Live Location Caption", result)
}

func TestGetMessageFromConversation(t *testing.T) {
	message := &proto.Message{
		Conversation: convert.ToPointer("Conversation Text"),
	}

	result := GetMessage(message)

	assert.Equal(t, "Conversation Text", result)
}

func TestGetMessageFromEmptyMessage(t *testing.T) {
	message := &proto.Message{}

	result := GetMessage(message)

	assert.Equal(t, "", result)
}
