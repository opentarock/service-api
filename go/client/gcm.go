package client

import "github.com/opentarock/service-api/go/proto_gcm"

type GcmParams struct {
	CollapseKey           string
	DelayWhileIdle        bool
	TimeToLive            uint64
	RestrictedPackageName string
}

type GcmClient interface {
	SendMessage(data string, params *GcmParams) (*proto_gcm.SendMessageResponse, error)
}
