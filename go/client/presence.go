package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_presence"
)

type PresenceClient interface {
	UpdateUserStatus(
		ctx context.Context,
		userId string,
		status proto_presence.UpdateUserStatusRequest_Status,
		device *proto_presence.Device) (*proto_presence.UpdateUserStatusResponse, error)

	GetUserDevices(
		ctx context.Context,
		userId string) (*proto_presence.GetUserDevicesResponse, error)
}
