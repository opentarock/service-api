package proto_user

import "github.com/opentarock/service-api/go/proto"

const (
	RegisterUserMessage     = 1
	AuthenticateUserMessage = 2
)

func (m *RegisterUser) GetMessageType() proto.Type {
	return RegisterUserMessage
}

func (m *AuthenticateUser) GetMessageType() proto.Type {
	return AuthenticateUserMessage
}

func NewInputError(name, errorMessage string) *RegisterResponse_InputError {
	return &RegisterResponse_InputError{
		Name:         &name,
		ErrorMessage: &errorMessage,
	}
}
