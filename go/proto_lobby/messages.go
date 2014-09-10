package proto_lobby

import (
	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	CreateRoomRequestMessage  = 0x0A0001
	CreateRoomResponseMessage = 0x0A0002

	ListRoomsRequestMessage  = 0x0A0010
	ListRoomsResponseMessage = 0x0A0011
)

func (m *CreateRoomRequest) GetMessageType() proto.Type {
	return CreateRoomRequestMessage
}

func (m *CreateRoomResponse) GetMessageType() proto.Type {
	return CreateRoomResponseMessage
}

func AsCreateRoomRequest(m *proto.Message) (*CreateRoomRequest, error) {
	const ty = "CreateRoomRequest"
	msg := CreateRoomRequest{}
	if m.Type != msg.GetMessageType() {
		return nil, proto.ErrorTypeMismatch(ty, msg.GetMessageType(), m.Type)
	}
	err := pbuf.Unmarshal(m.Data, &msg)
	if err != nil {
		return nil, proto.ErrorUnmarshalling(ty, err)
	}
	return &msg, nil
}

func AsCreateRoomResponse(m *proto.Message) (*CreateRoomResponse, error) {
	const ty = "CreateRoomResponse"
	msg := CreateRoomResponse{}
	if m.Type != msg.GetMessageType() {
		return nil, proto.ErrorTypeMismatch(ty, msg.GetMessageType(), m.Type)
	}
	err := pbuf.Unmarshal(m.Data, &msg)
	if err != nil {
		return nil, proto.ErrorUnmarshalling(ty, err)
	}
	return &msg, nil
}

func (m *ListRoomsRequest) GetMessageType() proto.Type {
	return ListRoomsRequestMessage
}

func (m *ListRoomsResponse) GetMessageType() proto.Type {
	return ListRoomsResponseMessage
}

func AsListRoomsRequest(m *proto.Message) (*ListRoomsRequest, error) {
	const ty = "ListRoomsRequest"
	msg := ListRoomsRequest{}
	if m.Type != msg.GetMessageType() {
		return nil, proto.ErrorTypeMismatch(ty, msg.GetMessageType(), m.Type)
	}
	err := pbuf.Unmarshal(m.Data, &msg)
	if err != nil {
		return nil, proto.ErrorUnmarshalling(ty, err)
	}
	return &msg, nil
}

func AsListRoomsResponse(m *proto.Message) (*ListRoomsResponse, error) {
	const ty = "ListRoomsResponse"
	msg := ListRoomsResponse{}
	if m.Type != msg.GetMessageType() {
		return nil, proto.ErrorTypeMismatch(ty, msg.GetMessageType(), m.Type)
	}
	err := pbuf.Unmarshal(m.Data, &msg)
	if err != nil {
		return nil, proto.ErrorUnmarshalling(ty, err)
	}
	return &msg, nil
}
