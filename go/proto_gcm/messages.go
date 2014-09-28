package proto_gcm

import "github.com/opentarock/service-api/go/proto"

const (
	SendMessageRequestMessage  = 0x110001
	SendMessageResponseMessage = 0x110002
)

func (m *SendMessageRequest) GetMessageType() proto.Type {
	return SendMessageRequestMessage
}

func (m *SendMessageResponse) GetMessageType() proto.Type {
	return SendMessageResponseMessage
}
