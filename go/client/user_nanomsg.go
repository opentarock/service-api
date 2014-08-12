package client

import (
	"log"
	"time"

	proto "code.google.com/p/gogoprotobuf/proto"
	nmsg "github.com/op/go-nanomsg"

	"github.com/opentarock/service-api/go/proto_user"
)

// UserClientNanomsg is an implementation of UserClient using nanomsg for message
// transport and protobuf for message serialization.
type UserClientNanomsg struct {
	userServiceSocket *nmsg.ReqSocket
}

func NewUserClientNanomsg() (*UserClientNanomsg, error) {
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
	endpoint, err := socket.Connect("tcp://127.0.0.1:6001")
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)

	return &UserClientNanomsg{
		userServiceSocket: socket,
	}, nil
}

func (s *UserClientNanomsg) RegisterUser(
	user *proto_user.User, redirectURI string) (*proto_user.RegisterResponse, error) {

	registerUser := &proto_user.RegisterUser{
		User:        user,
		RedirectUri: proto.String(redirectURI),
	}
	data, err := proto.Marshal(registerUser)
	if err != nil {
		log.Fatalf("Error marshaling RegisterUser: %s", err)
	}
	err = s.sendMsg(registerUser.GetMessageId(), data)
	if err != nil {
		return nil, err
	}
	responseData, err := s.userServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	response := &proto_user.RegisterResponse{}
	err = proto.Unmarshal(responseData, response)
	if err != nil {
		log.Fatalf("Error unmarshaling RegisterResponse: %s", err)
	}
	return response, nil
}

func (s *UserClientNanomsg) AuthenticateUser(email, password string) (*proto_user.AuthenticateResult, error) {

	authUser := &proto_user.AuthenticateUser{
		Email:    proto.String(email),
		Password: proto.String(password),
	}
	data, err := proto.Marshal(authUser)
	if err != nil {
		log.Fatalf("Error marshaling RegisterUser: %s", err)
	}
	err = s.sendMsg(authUser.GetMessageId(), data)
	if err != nil {
		return nil, err
	}
	responseData, err := s.userServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	response := &proto_user.AuthenticateResult{}
	err = proto.Unmarshal(responseData, response)
	if err != nil {
		log.Fatalf("Error unmarshaling RegisterResponse: %s", err)
	}
	return response, nil
}

func (s *UserClientNanomsg) sendMsg(messageId int, data []byte) error {
	prefixedData := make([]byte, 1, 1+len(data))
	prefixedData[0] = byte(messageId)
	prefixedData = append(prefixedData, data...)
	_, err := s.userServiceSocket.Send(prefixedData, 0)
	return err
}

// Close closes all the sockets and cleans up all the resources associated with
// this client.
// This method might block until all the resources are properly discarded.
func (s *UserClientNanomsg) Close() {
	s.userServiceSocket.Close()
}
