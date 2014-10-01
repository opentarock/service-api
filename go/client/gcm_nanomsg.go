package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_gcm"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
)

type GcmClientNanomsg struct {
	client *clientutil.ReqClient
}

func NewGcmClientNanomsg() *GcmClientNanomsg {
	return &GcmClientNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (c *GcmClientNanomsg) SendMessage(
	ctx context.Context,
	registrationIds []string,
	data string,
	params *proto_gcm.Parameters) (*proto_gcm.SendMessageResponse, error) {

	request := proto_gcm.SendMessageRequest{
		RegistrationIds: registrationIds,
		Params:          params,
	}
	if data != "" {
		request.Data = &data
	}

	var response proto_gcm.SendMessageResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(c.client, &request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *GcmClientNanomsg) Connect(address string) error {
	return c.client.Connect(address)
}

func (c *GcmClientNanomsg) Close() error {
	return c.client.Close()
}
