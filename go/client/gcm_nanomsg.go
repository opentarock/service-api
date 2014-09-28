package client

import "github.com/opentarock/service-api/go/proto_gcm"

type GcmClientNanomsg struct {
	client *ReqClient
}

func NewGcmClientNanomsg() *GcmClientNanomsg {
	return &GcmClientNanomsg{
		client: NewReqClient(),
	}
}

func (c *GcmClientNanomsg) SendMessage(data string, params *GcmParams) (*proto_gcm.SendMessageResponse, error) {

	request := proto_gcm.SendMessageRequest{
		Data: &data,
	}
	if params.CollapseKey != "" {
		request.CollapseKey = &params.CollapseKey
	}
	if params.DelayWhileIdle {
		request.DelayWhileIdle = &params.DelayWhileIdle
	}
	if params.TimeToLive != 0 {
		request.TimeToLive = &params.TimeToLive
	}
	if params.RestrictedPackageName != "" {
		request.RestrictedPackageName = &params.RestrictedPackageName
	}

	responseMsg, err := c.client.Request(&request)
	if err != nil {
		return nil, err
	}

	var response proto_gcm.SendMessageResponse
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
