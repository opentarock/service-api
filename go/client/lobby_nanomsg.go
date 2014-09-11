package client

import (
	"log"
	"time"

	pbuf "code.google.com/p/gogoprotobuf/proto"
	nmsg "github.com/op/go-nanomsg"

	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_headers"
	"github.com/opentarock/service-api/go/proto_lobby"
)

// UserClientNanomsg is an implementation of UserClient using nanomsg for message
// transport and protobuf for message serialization.
type LobbyClientNanomsg struct {
	lobbyServiceSocket *nmsg.ReqSocket
}

func NewLobbyClientNanomsg() (*LobbyClientNanomsg, error) {
	socket, err := nmsg.NewReqSocket()
	if err != nil {
		return nil, err
	}
	const timeout = 1 * time.Second
	err = socket.SetSendTimeout(timeout)
	err = socket.SetRecvTimeout(timeout)
	if err != nil {
		return nil, err
	}

	endpoint, err := socket.Connect("tcp://127.0.0.1:7001")
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)

	return &LobbyClientNanomsg{
		lobbyServiceSocket: socket,
	}, nil
}

func (s *LobbyClientNanomsg) CreateRoom(
	auth *proto_headers.AuthorizationHeader,
	name string,
	options *proto_lobby.RoomOptions) (*proto_lobby.CreateRoomResponse, error) {

	request := &proto_lobby.CreateRoomRequest{
		Name:    &name,
		Options: options,
	}

	msg, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	err = msg.Header.Marshal(auth)
	if err != nil {
		return nil, err
	}

	err = s.sendMsg(msg)
	if err != nil {
		return nil, err
	}
	responseData, err := s.lobbyServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	responseMsg, err := proto.Parse(responseData)
	if err != nil {
		return nil, err
	}
	var response proto_lobby.CreateRoomResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) ListRooms() (*proto_lobby.ListRoomsResponse, error) {

	request := &proto_lobby.ListRoomsRequest{}

	data, err := pbuf.Marshal(request)
	if err != nil {
		return nil, err
	}
	msg := proto.NewMessage(proto_lobby.ListRoomsRequestMessage, data)

	err = s.sendMsg(msg)
	if err != nil {
		return nil, err
	}
	responseData, err := s.lobbyServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	responseMsg, err := proto.Parse(responseData)
	if err != nil {
		return nil, err
	}
	var response proto_lobby.ListRoomsResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		log.Println(err)
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) sendMsg(msg *proto.Message) error {
	packed, err := msg.Pack()
	if err != nil {
		return err
	}
	_, err = s.lobbyServiceSocket.Send(packed, 0)
	return err
}

// Close closes all the sockets and cleans up all the resources associated with
// this client.
// This method might block until all the resources are properly discarded.
func (s *LobbyClientNanomsg) Close() {
	s.lobbyServiceSocket.Close()
}
