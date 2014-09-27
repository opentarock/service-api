package proto_presence

import "github.com/opentarock/service-api/go/proto"

const (
	UpdateUserStatusRequestMessage  = 0x080001
	UpdateUserStatusResponseMessage = 0x080002

	GetUserDevicesRequestMessage  = 0x080101
	GetUserDevicesResponseMessage = 0x080102
)

func (m *UpdateUserStatusRequest) GetMessageType() proto.Type {
	return UpdateUserStatusRequestMessage
}

func (m *UpdateUserStatusResponse) GetMessageType() proto.Type {
	return UpdateUserStatusResponseMessage
}

func (m *GetUserDevicesRequest) GetMessageType() proto.Type {
	return GetUserDevicesRequestMessage
}

func (m *GetUserDevicesResponse) GetMessageType() proto.Type {
	return GetUserDevicesResponseMessage
}
