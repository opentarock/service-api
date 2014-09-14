package client

import (
	"github.com/opentarock/service-api/go/proto_headers"
	"github.com/opentarock/service-api/go/proto_lobby"
)

type LobbyClient interface {
	CreateRoom(
		auth *proto_headers.AuthorizationHeader,
		name string,
		options *proto_lobby.RoomOptions) (*proto_lobby.CreateRoomResponse, error)

	JoinRoom(
		auth *proto_headers.AuthorizationHeader,
		roomId string) (*proto_lobby.JoinRoomResponse, error)

	LeaveRoom(
		auth *proto_headers.AuthorizationHeader) (*proto_lobby.LeaveRoomResponse, error)

	ListRooms(
		auth *proto_headers.AuthorizationHeader) (*proto_lobby.ListRoomsResponse, error)
}
