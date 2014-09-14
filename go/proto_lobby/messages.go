package proto_lobby

import "github.com/opentarock/service-api/go/proto"

const (
	CreateRoomRequestMessage  = 0x0A0001
	CreateRoomResponseMessage = 0x0A0002

	JoinRoomRequestMessage  = 0x0A0010
	JoinRoomResponseMessage = 0x0A0011

	LeaveRoomRequestMessage  = 0x0A0020
	LeaveRoomResponseMessage = 0x0A0021

	ListRoomsRequestMessage  = 0x0A00A0
	ListRoomsResponseMessage = 0x0A00A1
)

func (m *CreateRoomRequest) GetMessageType() proto.Type {
	return CreateRoomRequestMessage
}

func (m *CreateRoomResponse) GetMessageType() proto.Type {
	return CreateRoomResponseMessage
}

func (m *JoinRoomRequest) GetMessageType() proto.Type {
	return JoinRoomRequestMessage
}

func (m *JoinRoomResponse) GetMessageType() proto.Type {
	return JoinRoomResponseMessage
}

func (m *LeaveRoomRequest) GetMessageType() proto.Type {
	return LeaveRoomRequestMessage
}

func (m *LeaveRoomResponse) GetMessageType() proto.Type {
	return LeaveRoomResponseMessage
}

func (m *ListRoomsRequest) GetMessageType() proto.Type {
	return ListRoomsRequestMessage
}

func (m *ListRoomsResponse) GetMessageType() proto.Type {
	return ListRoomsResponseMessage
}
