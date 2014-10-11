package client

import (
	"code.google.com/p/go.net/context"

	"github.com/opentarock/service-api/go/proto_msghandler"
)

type MsgHandlerClient interface {
	RouteMessage(
		ctx context.Context,
		data string) (*proto_msghandler.RouteMessageResponse, error)
}
