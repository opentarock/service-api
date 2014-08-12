package client

import "github.com/opentarock/service-api/go/proto_user"

type UserClient interface {
	RegisterUser(user *proto_user.User, redirectURI string) (*proto_user.RegisterResponse, error)
	AuthenticateUser(email, password string) (*proto_user.AuthenticateResult, error)
}
