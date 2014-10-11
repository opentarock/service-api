package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_msghandler"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
)

type MsgHandlerClientNanomsg struct {
	client *clientutil.ReqClient
}

func NewMsgHandlerClientNanomsg() *MsgHandlerClientNanomsg {
	return &MsgHandlerClientNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (c *MsgHandlerClientNanomsg) RouteMessage(
	ctx context.Context,
	data string) (*proto_msghandler.RouteMessageResponse, error) {

	request := proto_msghandler.RouteMessageRequest{
		Data: &data,
	}

	var response proto_msghandler.RouteMessageResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, &request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *MsgHandlerClientNanomsg) Connect(address string) error {
	return c.client.Connect(address)
}

func (c *MsgHandlerClientNanomsg) Close() error {
	return c.client.Close()
}
