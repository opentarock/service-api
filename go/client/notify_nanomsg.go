package client

import (
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_notify"
)

type NotifyClientNanomsg struct {
	client *ReqClient
}

func NewNotifyClientNanomsg() *NotifyClientNanomsg {
	return &NotifyClientNanomsg{
		client: NewReqClient(),
	}
}

func (c *NotifyClientNanomsg) MessageUser(
	msg proto.ProtobufMessage, users ...uint64) (*proto_notify.MessageUsersResponse, error) {

	header := proto_notify.MessageUsersHeader{
		UserIds: users,
	}

	responseMsg, err := c.client.Request(msg, &header)
	if err != nil {
		return nil, err
	}

	var response proto_notify.MessageUsersResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *NotifyClientNanomsg) Connect(address string) error {
	return c.client.Connect(address)
}

func (c *NotifyClientNanomsg) Close() error {
	return c.client.Close()
}