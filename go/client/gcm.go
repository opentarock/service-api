package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_gcm"
)

type GcmClientFactory interface {
	WithContext(ctx context.Context) GcmClient
}

type GcmClient interface {
	SendMessage(
		registrationIds []string,
		data string,
		params *proto_gcm.Parameters) (*proto_gcm.SendMessageResponse, error)

	Client
}
