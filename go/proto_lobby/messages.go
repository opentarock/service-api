package proto_lobby

import "github.com/opentarock/service-api/go/proto"

const (
	CreateRoomRequestMessage  = 0x0A0001
	CreateRoomResponseMessage = 0x0A0002

	JoinRoomRequestMessage  = 0x0A0011
	JoinRoomResponseMessage = 0x0A0012
	JoinRoomEventMessage    = 0x0A0111

	LeaveRoomRequestMessage  = 0x0A0021
	LeaveRoomResponseMessage = 0x0A0022
	LeaveRoomEventMessage    = 0x0AF021

	ListRoomsRequestMessage  = 0x0A00A1
	ListRoomsResponseMessage = 0x0A00A2

	RoomInfoRequestMessage  = 0x0A00B1
	RoomInfoResponseMessage = 0x0A00B1

	StartGameRequestMessage  = 0x0A0101
	StartGameResponseMessage = 0x0A0102
	StartGameEventMessage    = 0x0AF101

	PlayerReadyRequestMessage  = 0x0A0201
	PlayerReadyResponseMessage = 0x0A0202
	PlayerReadyEventMessage    = 0x0AF201
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

func (m *JoinRoomEvent) GetMessageType() proto.Type {
	return JoinRoomEventMessage
}

func (m *LeaveRoomRequest) GetMessageType() proto.Type {
	return LeaveRoomRequestMessage
}

func (m *LeaveRoomResponse) GetMessageType() proto.Type {
	return LeaveRoomResponseMessage
}

func (m *LeaveRoomEvent) GetMessageType() proto.Type {
	return LeaveRoomEventMessage
}

func (m *ListRoomsRequest) GetMessageType() proto.Type {
	return ListRoomsRequestMessage
}

func (m *ListRoomsResponse) GetMessageType() proto.Type {
	return ListRoomsResponseMessage
}

func (m *RoomInfoRequest) GetMessageType() proto.Type {
	return RoomInfoRequestMessage
}

func (m *RoomInfoResponse) GetMessageType() proto.Type {
	return RoomInfoResponseMessage
}

func (m *StartGameRequest) GetMessageType() proto.Type {
	return StartGameRequestMessage
}

func (m *StartGameResponse) GetMessageType() proto.Type {
	return StartGameResponseMessage
}

func (m *StartGameEvent) GetMessageType() proto.Type {
	return StartGameEventMessage
}

func (m *PlayerReadyRequest) GetMessageType() proto.Type {
	return PlayerReadyRequestMessage
}

func (m *PlayerReadyResponse) GetMessageType() proto.Type {
	return PlayerReadyResponseMessage
}

func (m *PlayerReadyEvent) GetMessageType() proto.Type {
	return PlayerReadyEventMessage
}
