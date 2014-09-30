package client

import (
	"log"

	"github.com/opentarock/service-api/go/proto_headers"
	"github.com/opentarock/service-api/go/proto_lobby"
	"github.com/opentarock/service-api/go/util/clientutil"
)

// UserClientNanomsg is an implementation of UserClient using nanomsg for message
// transport and protobuf for message serialization.
type LobbyClientNanomsg struct {
	client *clientutil.ReqClient
}

func NewLobbyClientNanomsg() *LobbyClientNanomsg {
	return &LobbyClientNanomsg{
		client: clientutil.NewReqClient(),
	}
}

func (s *LobbyClientNanomsg) CreateRoom(
	auth *proto_headers.AuthorizationHeader,
	name string,
	options *proto_lobby.RoomOptions) (*proto_lobby.CreateRoomResponse, error) {

	request := &proto_lobby.CreateRoomRequest{
		Name:    &name,
		Options: options,
	}

	responseMsg, err := s.client.Request(request, auth)
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

func (s *LobbyClientNanomsg) JoinRoom(
	auth *proto_headers.AuthorizationHeader,
	roomId string) (*proto_lobby.JoinRoomResponse, error) {

	request := &proto_lobby.JoinRoomRequest{
		RoomId: &roomId,
	}

	responseMsg, err := s.client.Request(request, auth)
	if err != nil {
		return nil, err
	}

	var response proto_lobby.JoinRoomResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) LeaveRoom(
	auth *proto_headers.AuthorizationHeader) (*proto_lobby.LeaveRoomResponse, error) {

	request := &proto_lobby.LeaveRoomRequest{}

	responseMsg, err := s.client.Request(request, auth)
	if err != nil {
		return nil, err
	}

	var response proto_lobby.LeaveRoomResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) ListRooms(
	auth *proto_headers.AuthorizationHeader) (*proto_lobby.ListRoomsResponse, error) {

	request := &proto_lobby.ListRoomsRequest{}

	responseMsg, err := s.client.Request(request, auth)
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

func (s *LobbyClientNanomsg) RoomInfo(roomId string) (*proto_lobby.RoomInfoResponse, error) {
	request := &proto_lobby.RoomInfoRequest{
		RoomId: &roomId,
	}

	responseMsg, err := s.client.Request(request)
	if err != nil {
		return nil, err
	}

	var response proto_lobby.RoomInfoResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		log.Println(err)
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) StartGame(
	auth *proto_headers.AuthorizationHeader) (*proto_lobby.StartGameResponse, error) {

	request := &proto_lobby.StartGameRequest{}

	responseMsg, err := s.client.Request(request, auth)
	if err != nil {
		return nil, err
	}

	var response proto_lobby.StartGameResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		log.Println(err)
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) PlayerReady(
	auth *proto_headers.AuthorizationHeader,
	state string) (*proto_lobby.PlayerReadyResponse, error) {

	request := &proto_lobby.PlayerReadyRequest{
		State: &state,
	}

	responseMsg, err := s.client.Request(request, auth)
	if err != nil {
		return nil, err
	}

	var response proto_lobby.PlayerReadyResponse
	err = responseMsg.Unmarshal(&response)
	if err != nil {
		log.Println(err)
	}
	return &response, nil
}

func (s *LobbyClientNanomsg) Connect(address string) error {
	return s.client.Connect(address)
}

// Close closes all the sockets and cleans up all the resources associated with
// this client.
// This method might block until all the resources are properly discarded.
func (s *LobbyClientNanomsg) Close() error {
	return s.client.Close()
}
