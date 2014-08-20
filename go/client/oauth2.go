package client

import "github.com/opentarock/service-api/go/proto_oauth2"

type Oauth2Client interface {
	GetAccessToken(
		clientId, clientSecret string,
		request *proto_oauth2.AccessTokenRequest) (*proto_oauth2.AccessTokenResponse, error)
}
