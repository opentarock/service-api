package client

import (
	"code.google.com/p/go.net/context"

	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_notify"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
)

type NotifyClientNanomsg struct {
	client *clientutil.ReqClient
}

func NewNotifyClientNanomsg() *NotifyClientNanomsg {
	return &NotifyClientNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (c *NotifyClientNanomsg) MessageUsers(
	ctx context.Context,
	msg proto.ProtobufMessage,
	users ...string) (*proto_notify.MessageUsersResponse, error) {

	header := proto_notify.MessageUsersHeader{
		UserIds: users,
	}

	var response proto_notify.MessageUsersResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, msg, &response, &header)
	})
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
