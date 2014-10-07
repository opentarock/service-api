package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_presence"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
)

type PresenceClientNanomsg struct {
	client *clientutil.ReqClient
}

func NewPresenceClientNanomsg() *PresenceClientNanomsg {
	return &PresenceClientNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (c *PresenceClientNanomsg) UpdateUserStatus(
	ctx context.Context,
	userId string,
	status proto_presence.UpdateUserStatusRequest_Status,
	device *proto_presence.Device) (*proto_presence.UpdateUserStatusResponse, error) {

	request := &proto_presence.UpdateUserStatusRequest{
		UserId: &userId,
		Status: &status,
		Device: device,
	}

	var response proto_presence.UpdateUserStatusResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c PresenceClientNanomsg) GetUserDevices(
	ctx context.Context,
	userId string) (*proto_presence.GetUserDevicesResponse, error) {

	request := &proto_presence.GetUserDevicesRequest{
		UserId: &userId,
	}

	var response proto_presence.GetUserDevicesResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *PresenceClientNanomsg) Connect(address string) error {
	return c.client.Connect(address)
}

func (c *PresenceClientNanomsg) Close() error {
	return c.client.Close()
}
