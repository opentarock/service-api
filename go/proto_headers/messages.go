package proto_headers

import (
	"code.google.com/p/go-uuid/uuid"
	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	AuthorizationHeaderMessage      = 0x060101
	RequestCorrelationHeaderMessage = 0x060201
)

func (m *AuthorizationHeader) GetMessageType() proto.Type {
	return AuthorizationHeaderMessage
}

func NewAutorizationHeader(userId uint64, accessToken string) *AuthorizationHeader {
	return &AuthorizationHeader{
		UserId:      &userId,
		AccessToken: &accessToken,
	}
}

func (h *RequestCorrelationHeader) GetMessageType() proto.Type {
	return RequestCorrelationHeaderMessage
}

func NewRequestCorrelationHeader() *RequestCorrelationHeader {
	return &RequestCorrelationHeader{
		CorrelationId: pbuf.String(uuid.New()),
	}
}
