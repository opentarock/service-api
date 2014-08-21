package proto_user

import "code.google.com/p/gogoprotobuf/proto"

const (
	RegisterUserMessage     = 1
	AuthenticateUserMessage = 2
)

func (m *RegisterUser) GetMessageId() int {
	return RegisterUserMessage
}

func (m *AuthenticateUser) GetMessageId() int {
	return AuthenticateUserMessage
}

func NewInputError(name, errorMessage string) *RegisterResponse_InputError {
	return &RegisterResponse_InputError{
		Name:         proto.String(name),
		ErrorMessage: proto.String(errorMessage),
	}
}
