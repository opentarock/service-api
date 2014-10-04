package proto_notify

import (
	"encoding/json"

	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	MessageUsersHeaderMessage   = 0x090101
	MessageUsersResponseMessage = 0x090102

	TextMessageType = 0x090111
)

func (m *MessageUsersHeader) GetMessageType() proto.Type {
	return MessageUsersHeaderMessage
}

func (m *MessageUsersResponse) GetMessageType() proto.Type {
	return MessageUsersResponseMessage
}

func (m *TextMessage) GetMessageType() proto.Type {
	return TextMessageType
}

func NewJsonMessage(j map[string]interface{}) (*TextMessage, error) {
	enc, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return &TextMessage{
		Type: TextMessage_JSON.Enum(),
		Data: pbuf.String(string(enc)),
	}, nil
}
