package client

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto_lobby"
)

type LobbyClient interface {
	CreateRoom(
		ctx context.Context,
		name string,
		options *proto_lobby.RoomOptions) (*proto_lobby.CreateRoomResponse, error)

	JoinRoom(
		ctx context.Context,
		roomId string) (*proto_lobby.JoinRoomResponse, error)

	LeaveRoom(ctx context.Context) (*proto_lobby.LeaveRoomResponse, error)

	ListRooms(ctx context.Context) (*proto_lobby.ListRoomsResponse, error)

	RoomInfo(roomId string) (*proto_lobby.RoomInfoResponse, error)

	StartGame(ctx context.Context) (*proto_lobby.StartGameResponse, error)

	PlayerReady(
		ctx context.Context,
		state string) (*proto_lobby.PlayerReadyResponse, error)
}
