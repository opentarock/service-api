package proto_oauth2

import "github.com/opentarock/service-api/go/proto"

const (
	AuthorizationRequestMessage      = 1
	AccessTokenAuthenticationMessage = 2
	ValidateMessage                  = 3
)

func (m *AuthorizationRequest) GetMessageType() proto.Type {
	return AuthorizationRequestMessage
}

func (m *AccessTokenAuthentication) GetMessageType() proto.Type {
	return AccessTokenAuthenticationMessage
}

func (m *ValidateTokenRequest) GetMessageType() proto.Type {
	return ValidateMessage
}
