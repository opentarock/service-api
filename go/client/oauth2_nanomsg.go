package client

import (
	"log"
	"time"

	"code.google.com/p/goprotobuf/proto"

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
	const timeout = 100 * time.Millisecond
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

	accessTokenAuthentication := proto_oauth2.AccessTokenAuthentication{
		Client: &proto_oauth2.Client{
			Id:     proto.String(clientId),
			Secret: proto.String(clientSecret),
		},
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
	response := proto_oauth2.AccessTokenResponse{}
	err = proto.Unmarshal(responseData, &response)
	if err != nil {
		log.Fatalf("Error unmarshalling AccessTokenResponse: %s:", err)
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
