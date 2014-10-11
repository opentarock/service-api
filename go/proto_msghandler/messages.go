package proto_msghandler

import "github.com/opentarock/service-api/go/proto"

const (
	RouteMessageRequestType  = 0x120001
	RouteMessageResponseType = 0x120002
)

func (m *RouteMessageRequest) GetMessageType() proto.Type {
	return RouteMessageRequestType
}

func (m *RouteMessageResponse) GetMessageType() proto.Type {
	return RouteMessageResponseType
}
