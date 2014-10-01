package client

import (
	"github.com/opentarock/service-api/go/proto_presence"
	"github.com/opentarock/service-api/go/util/clientutil"
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
	userId string,
	status proto_presence.UpdateUserStatusRequest_Status,
	device *proto_presence.Device) (*proto_presence.UpdateUserStatusResponse, error) {

	request := &proto_presence.UpdateUserStatusRequest{
		UserId: &userId,
		Status: &status,
		Device: device,
	}

	responseMsg, err := c.client.Request(request)
	if err != nil {
		return nil, err
	}

	var response proto_presence.UpdateUserStatusResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c PresenceClientNanomsg) GetUserDevices(userId string) (*proto_presence.GetUserDevicesResponse, error) {
	request := &proto_presence.GetUserDevicesRequest{
		UserId: &userId,
	}

	responseMsg, err := c.client.Request(request)
	if err != nil {
		return nil, err
	}

	var response proto_presence.GetUserDevicesResponse
	err = responseMsg.Unmarshal(&response)
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
