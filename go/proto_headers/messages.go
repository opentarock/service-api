package proto_headers

import (
	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	AuthorizationHeaderMessage = 0x060101
)

func (m *AuthorizationHeader) GetMessageType() proto.Type {
	return AuthorizationHeaderMessage
}

func GetAuthorizationHeader(m *proto.Message) (*AuthorizationHeader, error) {
	authorizationHeader := AuthorizationHeader{}
	data := m.Header.Get(authorizationHeader.GetMessageType())
	if data == nil {
		return nil, nil
	}
	err := pbuf.Unmarshal(data, &authorizationHeader)
	if err != nil {
		return nil, proto.ErrorUnmarshalling("AuthorizationHeader", err)
	}
	return &authorizationHeader, nil
}
