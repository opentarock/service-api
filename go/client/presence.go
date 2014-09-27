package client

import "github.com/opentarock/service-api/go/proto_presence"

type PresenceClient interface {
	UpdateUserStatus(
		userId string,
		status proto_presence.UpdateUserStatusRequest_Status,
		device *proto_presence.Device) (*proto_presence.UpdateUserStatusResponse, error)

	GetUserDevices(userId string) (*proto_presence.GetUserDevicesResponse, error)
}
