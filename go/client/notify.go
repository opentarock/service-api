package client

import (
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_notify"
)

type NotifyClient interface {
	MessageUsers(msg proto.ProtobufMessage, users ...uint64) (*proto_notify.MessageUsersResponse, error)
}
