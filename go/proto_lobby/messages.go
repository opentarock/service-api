package proto_lobby

import "github.com/opentarock/service-api/go/proto"

const (
	CreateRoomRequestMessage  = 0x0A0001
	CreateRoomResponseMessage = 0x0A0002

	JoinRoomRequestMessage  = 0x0A0011
	JoinRoomResponseMessage = 0x0A0012

	ListRoomsRequestMessage  = 0x0A0020
	ListRoomsResponseMessage = 0x0A0021
)

func (m *CreateRoomRequest) GetMessageType() proto.Type {
	return CreateRoomRequestMessage
}

func (m *CreateRoomResponse) GetMessageType() proto.Type {
	return CreateRoomResponseMessage
}

func (m *ListRoomsRequest) GetMessageType() proto.Type {
	return ListRoomsRequestMessage
}

func (m *ListRoomsResponse) GetMessageType() proto.Type {
	return ListRoomsResponseMessage
}
