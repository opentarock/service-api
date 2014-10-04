package client

import (
	"code.google.com/p/go.net/context"

	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_notify"
)

type NotifyClient interface {
	MessageUsers(
		ctx context.Context,
		msg proto.ProtobufMessage,
		users ...uint64) (*proto_notify.MessageUsersResponse, error)
}
