package proto_notify

import "github.com/opentarock/service-api/go/proto"

const (
	MessageUsersHeaderMessage   = 0x090001
	MessageUsersResponseMessage = 0x090002
)

func (m *MessageUsersHeader) GetMessageType() proto.Type {
	return MessageUsersHeaderMessage
}

func (m *MessageUsersResponse) GetMessageType() proto.Type {
	return MessageUsersResponseMessage
}
