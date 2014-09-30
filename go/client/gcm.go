package client

import "github.com/opentarock/service-api/go/proto_gcm"

type GcmClient interface {
	SendMessage(
		registrationIds []string,
		data string,
		params *proto_gcm.Parameters) (*proto_gcm.SendMessageResponse, error)

	Client
}
