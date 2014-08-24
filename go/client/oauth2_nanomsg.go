package client

import (
	"errors"
	"log"
	"time"

	"code.google.com/p/gogoprotobuf/proto"

	nmsg "github.com/op/go-nanomsg"
	"github.com/opentarock/service-api/go/proto_oauth2"
)

// Oauth2ClientNanomsg is and implementation of Oauth2Client using nanomsg for
// message transport and protobuf for message serialization.
type Oauth2ClientNanomsg struct {
	oauth2ServiceSocket *nmsg.ReqSocket
}

// TODO: remove duplication with UserClientNanomsg
func NewOauth2ClientNanomsg() (*Oauth2ClientNanomsg, error) {
	socket, err := nmsg.NewReqSocket()
	if err != nil {
		return nil, err
	}
	// Timeout is set because we can't wait for the messages forever to keep
	// the frontend responsive.
	const timeout = 1000 * time.Millisecond
	err = socket.SetSendTimeout(timeout)
	err = socket.SetRecvTimeout(timeout)
	if err != nil {
		return nil, err
	}

	// TODO: Make address and port a parameter
	endpoint, err := socket.Connect("tcp://127.0.0.1:6002")
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)

	return &Oauth2ClientNanomsg{
		oauth2ServiceSocket: socket,
	}, nil
}

func (s *Oauth2ClientNanomsg) GetAccessToken(
	clientId, clientSecret string, request *proto_oauth2.AccessTokenRequest) (*proto_oauth2.AccessTokenResponse, error) {

	var client *proto_oauth2.Client
	if clientId != "" || clientSecret != "" {
		client = &proto_oauth2.Client{}
		if clientId != "" {
			client.Id = &clientId
		}
		if clientSecret != "" {
			client.Secret = &clientSecret
		}
	}

	accessTokenAuthentication := proto_oauth2.AccessTokenAuthentication{
		Client:  client,
		Request: request,
	}
	data, err := proto.Marshal(&accessTokenAuthentication)
	if err != nil {
		log.Fatalf("Error marshalling AccessTokenAuthentication: %s", err)
	}
	err = s.sendMsg(accessTokenAuthentication.GetMessageId(), data)
	if err != nil {
		return nil, err
	}
	responseData, err := s.oauth2ServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	if len(responseData) == 0 {
		return nil, errors.New("Empty response")
	}
	response := proto_oauth2.AccessTokenResponse{}
	err = proto.Unmarshal(responseData, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling AccessTokenResponse: %s:", err)
	}
	return &response, nil
}

func (s *Oauth2ClientNanomsg) ValidateToken(accessToken string) (*proto_oauth2.ValidateTokenResponse, error) {
	request := proto_oauth2.ValidateTokenRequest{
		AccessToken: &accessToken,
	}

	data, err := proto.Marshal(&request)
	if err != nil {
		log.Fatalf("Error marshalling ValidateTokenRequest: %s", err)
	}
	err = s.sendMsg(request.GetMessageId(), data)
	if err != nil {
		return nil, err
	}
	responseData, err := s.recvMsg()
	if err != nil {
		return nil, err
	}
	response := proto_oauth2.ValidateTokenResponse{}
	err = proto.Unmarshal(responseData, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling ValidateTokenResponse: %s:", err)
	}
	return &response, nil
}

func (s *Oauth2ClientNanomsg) sendMsg(messageId int, data []byte) error {
	prefixedData := make([]byte, 1, 1+len(data))
	prefixedData[0] = byte(messageId)
	prefixedData = append(prefixedData, data...)
	_, err := s.oauth2ServiceSocket.Send(prefixedData, 0)
	return err
}

func (s *Oauth2ClientNanomsg) recvMsg() ([]byte, error) {
	responseData, err := s.oauth2ServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	if len(responseData) == 0 {
		return nil, errors.New("Empty response")
	}
	return responseData, nil
}
