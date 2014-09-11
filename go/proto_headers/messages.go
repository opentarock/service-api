package proto_headers

import "github.com/opentarock/service-api/go/proto"

const (
	AuthorizationHeaderMessage = 0x060101
)

func (m *AuthorizationHeader) GetMessageType() proto.Type {
	return AuthorizationHeaderMessage
}
