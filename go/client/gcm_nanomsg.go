package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_gcm"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
)

type GcmClientFactoryNanomsg struct {
	client *clientutil.ReqClient
}

func NewGcmClientFactoryNanomsg() *GcmClientFactoryNanomsg {
	return &GcmClientFactoryNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (f *GcmClientFactoryNanomsg) WithContext(ctx context.Context) GcmClient {
	return NewGcmClientNanomsg(ctx, f.client)
}

type GcmClientNanomsg struct {
	ctx    context.Context
	client *clientutil.ReqClient
}

func NewGcmClientNanomsg(ctx context.Context, client *clientutil.ReqClient) *GcmClientNanomsg {
	return &GcmClientNanomsg{
		ctx:    ctx,
		client: client,
	}
}

func (c *GcmClientNanomsg) SendMessage(
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
	err := contextutil.Do(c.ctx, func() error {
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
