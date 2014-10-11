package client

import (
	"code.google.com/p/go.net/context"

	"github.com/opentarock/service-api/go/proto_lobby"
	"github.com/opentarock/service-api/go/util/clientutil"
	"github.com/opentarock/service-api/go/util/contextutil"
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

func (c *LobbyClientNanomsg) CreateRoom(
	ctx context.Context,
	name string,
	options *proto_lobby.RoomOptions) (*proto_lobby.CreateRoomResponse, error) {

	request := &proto_lobby.CreateRoomRequest{
		Name:    &name,
		Options: options,
	}

	var response proto_lobby.CreateRoomResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) JoinRoom(
	ctx context.Context,
	roomId string) (*proto_lobby.JoinRoomResponse, error) {

	request := &proto_lobby.JoinRoomRequest{
		RoomId: &roomId,
	}

	var response proto_lobby.JoinRoomResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) LeaveRoom(
	ctx context.Context) (*proto_lobby.LeaveRoomResponse, error) {

	request := &proto_lobby.LeaveRoomRequest{}

	var response proto_lobby.LeaveRoomResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) ListRooms(
	ctx context.Context) (*proto_lobby.ListRoomsResponse, error) {

	request := &proto_lobby.ListRoomsRequest{}

	var response proto_lobby.ListRoomsResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) RoomInfo(
	ctx context.Context,
	roomId string) (*proto_lobby.RoomInfoResponse, error) {

	request := &proto_lobby.RoomInfoRequest{
		RoomId: &roomId,
	}

	var response proto_lobby.RoomInfoResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) StartGame(
	ctx context.Context) (*proto_lobby.StartGameResponse, error) {

	request := &proto_lobby.StartGameRequest{}

	var response proto_lobby.StartGameResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *LobbyClientNanomsg) PlayerReady(
	ctx context.Context,
	state string) (*proto_lobby.PlayerReadyResponse, error) {

	request := &proto_lobby.PlayerReadyRequest{
		State: &state,
	}

	var response proto_lobby.PlayerReadyResponse
	err := contextutil.Do(ctx, func() error {
		return clientutil.DoRequest(ctx, c.client, request, &response)
	})
	if err != nil {
		return nil, err
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
