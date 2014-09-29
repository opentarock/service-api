package client

import (
	"github.com/opentarock/service-api/go/proto_gcm"
	"github.com/opentarock/service-api/go/util/clientutil"
)

type GcmClientNanomsg struct {
	client *ReqClient
}

func NewGcmClientNanomsg() *GcmClientNanomsg {
	return &GcmClientNanomsg{
		client: NewReqClient(),
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

	responseMsg, err := c.client.Request(&request)
	if err != nil {
		return nil, err
	}

	var response proto_gcm.SendMessageResponse
	err = clientutil.TryDecodeError(responseMsg, &response)
	if err != nil {
		return nil, err
	}

	err = responseMsg.Unmarshal(&response)
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
