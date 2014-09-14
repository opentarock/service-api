package proto_notify

import "github.com/opentarock/service-api/go/proto"

const (
	MessageUsersHeaderMessage   = 0x090000
	MessageUsersResponseMessage = 0x090000
)

func (m *MessageUsersHeader) GetMessageType() proto.Type {
	return MessageUsersHeaderMessage
}

func (m *MessageUsersResponse) GetMessageType() proto.Type {
	return MessageUsersResponseMessage
}
