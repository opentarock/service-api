package proto_oauth2

const (
	AuthorizationRequestMessage      = 1
	AccessTokenAuthenticationMessage = 2
	ValidateMessage                  = 3
)

func (m *AuthorizationRequest) GetMessageId() int {
	return AuthorizationRequestMessage
}

func (m *AccessTokenAuthentication) GetMessageId() int {
	return AccessTokenAuthenticationMessage
}

func (m *ValidateTokenRequest) GetMessageId() int {
	return ValidateMessage
}
