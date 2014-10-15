// Code generated by protoc-gen-gogo.
// source: lobby.proto
// DO NOT EDIT!

/*
Package proto_lobby is a generated protocol buffer package.

It is generated from these files:
	lobby.proto

It has these top-level messages:
	Room
	RoomOptions
	CreateRoomRequest
	CreateRoomResponse
	JoinRoomRequest
	JoinRoomResponse
	JoinRoomEvent
	LeaveRoomRequest
	LeaveRoomResponse
	LeaveRoomEvent
	ListRoomsRequest
	ListRoomsResponse
	RoomInfoRequest
	RoomInfoResponse
	StartGameRequest
	StartGameResponse
	StartGameEvent
	PlayerReadyRequest
	PlayerReadyResponse
	PlayerReadyEvent
*/
package proto_lobby

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CreateRoomResponse_ErrorCode int32

const (
	CreateRoomResponse_ALREADY_IN_ROOM CreateRoomResponse_ErrorCode = 0
)

var CreateRoomResponse_ErrorCode_name = map[int32]string{
	0: "ALREADY_IN_ROOM",
}
var CreateRoomResponse_ErrorCode_value = map[string]int32{
	"ALREADY_IN_ROOM": 0,
}

func (x CreateRoomResponse_ErrorCode) Enum() *CreateRoomResponse_ErrorCode {
	p := new(CreateRoomResponse_ErrorCode)
	*p = x
	return p
}
func (x CreateRoomResponse_ErrorCode) String() string {
	return proto.EnumName(CreateRoomResponse_ErrorCode_name, int32(x))
}
func (x *CreateRoomResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CreateRoomResponse_ErrorCode_value, data, "CreateRoomResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = CreateRoomResponse_ErrorCode(value)
	return nil
}

type JoinRoomResponse_ErrorCode int32

const (
	JoinRoomResponse_ROOM_DOES_NOT_EXIST JoinRoomResponse_ErrorCode = 0
	JoinRoomResponse_ROOM_FULL           JoinRoomResponse_ErrorCode = 1
)

var JoinRoomResponse_ErrorCode_name = map[int32]string{
	0: "ROOM_DOES_NOT_EXIST",
	1: "ROOM_FULL",
}
var JoinRoomResponse_ErrorCode_value = map[string]int32{
	"ROOM_DOES_NOT_EXIST": 0,
	"ROOM_FULL":           1,
}

func (x JoinRoomResponse_ErrorCode) Enum() *JoinRoomResponse_ErrorCode {
	p := new(JoinRoomResponse_ErrorCode)
	*p = x
	return p
}
func (x JoinRoomResponse_ErrorCode) String() string {
	return proto.EnumName(JoinRoomResponse_ErrorCode_name, int32(x))
}
func (x *JoinRoomResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(JoinRoomResponse_ErrorCode_value, data, "JoinRoomResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = JoinRoomResponse_ErrorCode(value)
	return nil
}

type LeaveRoomResponse_ErrorCode int32

const (
	LeaveRoomResponse_NOT_IN_ROOM       LeaveRoomResponse_ErrorCode = 0
	LeaveRoomResponse_START_IN_PROGRESS LeaveRoomResponse_ErrorCode = 1
)

var LeaveRoomResponse_ErrorCode_name = map[int32]string{
	0: "NOT_IN_ROOM",
	1: "START_IN_PROGRESS",
}
var LeaveRoomResponse_ErrorCode_value = map[string]int32{
	"NOT_IN_ROOM":       0,
	"START_IN_PROGRESS": 1,
}

func (x LeaveRoomResponse_ErrorCode) Enum() *LeaveRoomResponse_ErrorCode {
	p := new(LeaveRoomResponse_ErrorCode)
	*p = x
	return p
}
func (x LeaveRoomResponse_ErrorCode) String() string {
	return proto.EnumName(LeaveRoomResponse_ErrorCode_name, int32(x))
}
func (x *LeaveRoomResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LeaveRoomResponse_ErrorCode_value, data, "LeaveRoomResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = LeaveRoomResponse_ErrorCode(value)
	return nil
}

type RoomInfoResponse_ErrorCode int32

const (
	RoomInfoResponse_ROOM_DOES_NOT_EXIST RoomInfoResponse_ErrorCode = 0
)

var RoomInfoResponse_ErrorCode_name = map[int32]string{
	0: "ROOM_DOES_NOT_EXIST",
}
var RoomInfoResponse_ErrorCode_value = map[string]int32{
	"ROOM_DOES_NOT_EXIST": 0,
}

func (x RoomInfoResponse_ErrorCode) Enum() *RoomInfoResponse_ErrorCode {
	p := new(RoomInfoResponse_ErrorCode)
	*p = x
	return p
}
func (x RoomInfoResponse_ErrorCode) String() string {
	return proto.EnumName(RoomInfoResponse_ErrorCode_name, int32(x))
}
func (x *RoomInfoResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RoomInfoResponse_ErrorCode_value, data, "RoomInfoResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = RoomInfoResponse_ErrorCode(value)
	return nil
}

type StartGameResponse_ErrorCode int32

const (
	StartGameResponse_NOT_IN_ROOM     StartGameResponse_ErrorCode = 0
	StartGameResponse_NOT_OWNER       StartGameResponse_ErrorCode = 1
	StartGameResponse_ALREADY_STARTED StartGameResponse_ErrorCode = 2
)

var StartGameResponse_ErrorCode_name = map[int32]string{
	0: "NOT_IN_ROOM",
	1: "NOT_OWNER",
	2: "ALREADY_STARTED",
}
var StartGameResponse_ErrorCode_value = map[string]int32{
	"NOT_IN_ROOM":     0,
	"NOT_OWNER":       1,
	"ALREADY_STARTED": 2,
}

func (x StartGameResponse_ErrorCode) Enum() *StartGameResponse_ErrorCode {
	p := new(StartGameResponse_ErrorCode)
	*p = x
	return p
}
func (x StartGameResponse_ErrorCode) String() string {
	return proto.EnumName(StartGameResponse_ErrorCode_name, int32(x))
}
func (x *StartGameResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StartGameResponse_ErrorCode_value, data, "StartGameResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = StartGameResponse_ErrorCode(value)
	return nil
}

type PlayerReadyResponse_ErrorCode int32

const (
	PlayerReadyResponse_NOT_IN_ROOM   PlayerReadyResponse_ErrorCode = 0
	PlayerReadyResponse_UNEXPECTED    PlayerReadyResponse_ErrorCode = 1
	PlayerReadyResponse_ALREADY_READY PlayerReadyResponse_ErrorCode = 2
	PlayerReadyResponse_INVALID_STATE PlayerReadyResponse_ErrorCode = 3
)

var PlayerReadyResponse_ErrorCode_name = map[int32]string{
	0: "NOT_IN_ROOM",
	1: "UNEXPECTED",
	2: "ALREADY_READY",
	3: "INVALID_STATE",
}
var PlayerReadyResponse_ErrorCode_value = map[string]int32{
	"NOT_IN_ROOM":   0,
	"UNEXPECTED":    1,
	"ALREADY_READY": 2,
	"INVALID_STATE": 3,
}

func (x PlayerReadyResponse_ErrorCode) Enum() *PlayerReadyResponse_ErrorCode {
	p := new(PlayerReadyResponse_ErrorCode)
	*p = x
	return p
}
func (x PlayerReadyResponse_ErrorCode) String() string {
	return proto.EnumName(PlayerReadyResponse_ErrorCode_name, int32(x))
}
func (x *PlayerReadyResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayerReadyResponse_ErrorCode_value, data, "PlayerReadyResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = PlayerReadyResponse_ErrorCode(value)
	return nil
}

type Room struct {
	Id               *string      `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Options          *RoomOptions `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	Owner            *string      `protobuf:"bytes,4,req,name=owner" json:"owner,omitempty"`
	Players          []string     `protobuf:"bytes,5,rep,name=players" json:"players,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}

func (m *Room) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Room) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Room) GetOptions() *RoomOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Room) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *Room) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

type RoomOptions struct {
	Private          *bool   `protobuf:"varint,1,opt,name=private,def=0" json:"private,omitempty"`
	Password         *string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Autostart        *bool   `protobuf:"varint,3,opt,name=autostart,def=0" json:"autostart,omitempty"`
	FillBots         *bool   `protobuf:"varint,4,opt,name=fill_bots,def=1" json:"fill_bots,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomOptions) Reset()         { *m = RoomOptions{} }
func (m *RoomOptions) String() string { return proto.CompactTextString(m) }
func (*RoomOptions) ProtoMessage()    {}

const Default_RoomOptions_Private bool = false
const Default_RoomOptions_Autostart bool = false
const Default_RoomOptions_FillBots bool = true

func (m *RoomOptions) GetPrivate() bool {
	if m != nil && m.Private != nil {
		return *m.Private
	}
	return Default_RoomOptions_Private
}

func (m *RoomOptions) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *RoomOptions) GetAutostart() bool {
	if m != nil && m.Autostart != nil {
		return *m.Autostart
	}
	return Default_RoomOptions_Autostart
}

func (m *RoomOptions) GetFillBots() bool {
	if m != nil && m.FillBots != nil {
		return *m.FillBots
	}
	return Default_RoomOptions_FillBots
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type CreateRoomRequest struct {
	Name             *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Options          *RoomOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}

func (m *CreateRoomRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateRoomRequest) GetOptions() *RoomOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateRoomResponse struct {
	Room             *Room                         `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
	ErrorCode        *CreateRoomResponse_ErrorCode `protobuf:"varint,2,opt,name=error_code,enum=proto_lobby.CreateRoomResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *CreateRoomResponse) Reset()         { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()    {}

func (m *CreateRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *CreateRoomResponse) GetErrorCode() CreateRoomResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return CreateRoomResponse_ALREADY_IN_ROOM
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type JoinRoomRequest struct {
	RoomId           *string `protobuf:"bytes,1,req,name=room_id" json:"room_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JoinRoomRequest) Reset()         { *m = JoinRoomRequest{} }
func (m *JoinRoomRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRequest) ProtoMessage()    {}

func (m *JoinRoomRequest) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

type JoinRoomResponse struct {
	Room             *Room                       `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
	ErrorCode        *JoinRoomResponse_ErrorCode `protobuf:"varint,2,opt,name=error_code,enum=proto_lobby.JoinRoomResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *JoinRoomResponse) Reset()         { *m = JoinRoomResponse{} }
func (m *JoinRoomResponse) String() string { return proto.CompactTextString(m) }
func (*JoinRoomResponse) ProtoMessage()    {}

func (m *JoinRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *JoinRoomResponse) GetErrorCode() JoinRoomResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return JoinRoomResponse_ROOM_DOES_NOT_EXIST
}

type JoinRoomEvent struct {
	Player           *string `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JoinRoomEvent) Reset()         { *m = JoinRoomEvent{} }
func (m *JoinRoomEvent) String() string { return proto.CompactTextString(m) }
func (*JoinRoomEvent) ProtoMessage()    {}

func (m *JoinRoomEvent) GetPlayer() string {
	if m != nil && m.Player != nil {
		return *m.Player
	}
	return ""
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type LeaveRoomRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *LeaveRoomRequest) Reset()         { *m = LeaveRoomRequest{} }
func (m *LeaveRoomRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomRequest) ProtoMessage()    {}

type LeaveRoomResponse struct {
	ErrorCode        *LeaveRoomResponse_ErrorCode `protobuf:"varint,1,opt,name=error_code,enum=proto_lobby.LeaveRoomResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *LeaveRoomResponse) Reset()         { *m = LeaveRoomResponse{} }
func (m *LeaveRoomResponse) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomResponse) ProtoMessage()    {}

func (m *LeaveRoomResponse) GetErrorCode() LeaveRoomResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return LeaveRoomResponse_NOT_IN_ROOM
}

type LeaveRoomEvent struct {
	PlayerId         *uint64 `protobuf:"varint,1,req,name=player_id" json:"player_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LeaveRoomEvent) Reset()         { *m = LeaveRoomEvent{} }
func (m *LeaveRoomEvent) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomEvent) ProtoMessage()    {}

func (m *LeaveRoomEvent) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type ListRoomsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListRoomsRequest) Reset()         { *m = ListRoomsRequest{} }
func (m *ListRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRoomsRequest) ProtoMessage()    {}

type ListRoomsResponse struct {
	Rooms            []*Room `protobuf:"bytes,1,rep,name=rooms" json:"rooms,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListRoomsResponse) Reset()         { *m = ListRoomsResponse{} }
func (m *ListRoomsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRoomsResponse) ProtoMessage()    {}

func (m *ListRoomsResponse) GetRooms() []*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type RoomInfoRequest struct {
	RoomId           *string `protobuf:"bytes,1,req,name=room_id" json:"room_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomInfoRequest) Reset()         { *m = RoomInfoRequest{} }
func (m *RoomInfoRequest) String() string { return proto.CompactTextString(m) }
func (*RoomInfoRequest) ProtoMessage()    {}

func (m *RoomInfoRequest) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

type RoomInfoResponse struct {
	ErrorCode        *RoomInfoResponse_ErrorCode `protobuf:"varint,1,opt,name=error_code,enum=proto_lobby.RoomInfoResponse_ErrorCode" json:"error_code,omitempty"`
	Room             *Room                       `protobuf:"bytes,2,opt,name=room" json:"room,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *RoomInfoResponse) Reset()         { *m = RoomInfoResponse{} }
func (m *RoomInfoResponse) String() string { return proto.CompactTextString(m) }
func (*RoomInfoResponse) ProtoMessage()    {}

func (m *RoomInfoResponse) GetErrorCode() RoomInfoResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return RoomInfoResponse_ROOM_DOES_NOT_EXIST
}

func (m *RoomInfoResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type StartGameRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StartGameRequest) Reset()         { *m = StartGameRequest{} }
func (m *StartGameRequest) String() string { return proto.CompactTextString(m) }
func (*StartGameRequest) ProtoMessage()    {}

type StartGameResponse struct {
	ErrorCode        *StartGameResponse_ErrorCode `protobuf:"varint,1,opt,name=error_code,enum=proto_lobby.StartGameResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *StartGameResponse) Reset()         { *m = StartGameResponse{} }
func (m *StartGameResponse) String() string { return proto.CompactTextString(m) }
func (*StartGameResponse) ProtoMessage()    {}

func (m *StartGameResponse) GetErrorCode() StartGameResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return StartGameResponse_NOT_IN_ROOM
}

type StartGameEvent struct {
	RoomId           *string `protobuf:"bytes,1,req,name=room_id" json:"room_id,omitempty"`
	State            *string `protobuf:"bytes,2,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StartGameEvent) Reset()         { *m = StartGameEvent{} }
func (m *StartGameEvent) String() string { return proto.CompactTextString(m) }
func (*StartGameEvent) ProtoMessage()    {}

func (m *StartGameEvent) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

func (m *StartGameEvent) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

// Requires proto_headers.AuthorizationHeader to be sent as header.
type PlayerReadyRequest struct {
	State            *string `protobuf:"bytes,1,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerReadyRequest) Reset()         { *m = PlayerReadyRequest{} }
func (m *PlayerReadyRequest) String() string { return proto.CompactTextString(m) }
func (*PlayerReadyRequest) ProtoMessage()    {}

func (m *PlayerReadyRequest) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

type PlayerReadyResponse struct {
	ErrorCode        *PlayerReadyResponse_ErrorCode `protobuf:"varint,1,opt,name=error_code,enum=proto_lobby.PlayerReadyResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *PlayerReadyResponse) Reset()         { *m = PlayerReadyResponse{} }
func (m *PlayerReadyResponse) String() string { return proto.CompactTextString(m) }
func (*PlayerReadyResponse) ProtoMessage()    {}

func (m *PlayerReadyResponse) GetErrorCode() PlayerReadyResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return PlayerReadyResponse_NOT_IN_ROOM
}

type PlayerReadyEvent struct {
	UserId           *string `protobuf:"bytes,1,req,name=user_id" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerReadyEvent) Reset()         { *m = PlayerReadyEvent{} }
func (m *PlayerReadyEvent) String() string { return proto.CompactTextString(m) }
func (*PlayerReadyEvent) ProtoMessage()    {}

func (m *PlayerReadyEvent) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto_lobby.CreateRoomResponse_ErrorCode", CreateRoomResponse_ErrorCode_name, CreateRoomResponse_ErrorCode_value)
	proto.RegisterEnum("proto_lobby.JoinRoomResponse_ErrorCode", JoinRoomResponse_ErrorCode_name, JoinRoomResponse_ErrorCode_value)
	proto.RegisterEnum("proto_lobby.LeaveRoomResponse_ErrorCode", LeaveRoomResponse_ErrorCode_name, LeaveRoomResponse_ErrorCode_value)
	proto.RegisterEnum("proto_lobby.RoomInfoResponse_ErrorCode", RoomInfoResponse_ErrorCode_name, RoomInfoResponse_ErrorCode_value)
	proto.RegisterEnum("proto_lobby.StartGameResponse_ErrorCode", StartGameResponse_ErrorCode_name, StartGameResponse_ErrorCode_value)
	proto.RegisterEnum("proto_lobby.PlayerReadyResponse_ErrorCode", PlayerReadyResponse_ErrorCode_name, PlayerReadyResponse_ErrorCode_value)
}
