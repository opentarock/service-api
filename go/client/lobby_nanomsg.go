package client

import (
	"log"
	"time"

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

	responseMsg, err := s.rpcCall(request, auth)
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

	responseMsg, err := s.rpcCall(request, auth)
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

	responseMsg, err := s.rpcCall(request, auth)
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

	responseMsg, err := s.rpcCall(request, auth)
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

	responseMsg, err := s.rpcCall(request)
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

	responseMsg, err := s.rpcCall(request, auth)
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

	responseMsg, err := s.rpcCall(request, auth)
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

func (s *LobbyClientNanomsg) rpcCall(
	request proto.ProtobufMessage,
	headers ...proto.ProtobufMessage) (*proto.Message, error) {

	msg, err := proto.MarshalHeaders(request, headers)
	if err != nil {
		return nil, err
	}

	err = s.sendMsg(msg)
	if err != nil {
		return nil, err
	}
	responseMsg, err := s.recvMsg()
	if err != nil {
		return nil, err
	}

	return responseMsg, nil
}

func (s *LobbyClientNanomsg) sendMsg(msg *proto.Message) error {
	packed, err := msg.Pack()
	if err != nil {
		return err
	}
	_, err = s.lobbyServiceSocket.Send(packed, 0)
	return err
}

func (s *LobbyClientNanomsg) recvMsg() (*proto.Message, error) {
	responseData, err := s.lobbyServiceSocket.Recv(0)
	if err != nil {
		return nil, err
	}
	responseMsg, err := proto.Parse(responseData)
	if err != nil {
		return nil, err
	}
	return responseMsg, err
}

// Close closes all the sockets and cleans up all the resources associated with
// this client.
// This method might block until all the resources are properly discarded.
func (s *LobbyClientNanomsg) Close() {
	s.lobbyServiceSocket.Close()
}
