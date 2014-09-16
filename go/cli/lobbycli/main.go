package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_headers"
	"github.com/opentarock/service-api/go/proto_lobby"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s ACTION [ARGS...]\n", os.Args[0])
		os.Exit(1)
	}
	client, err := client.NewLobbyClientNanomsg()
	exitError(err)
	switch os.Args[1] {
	case "create_room":
		userId, err := strconv.ParseUint(getArg(2), 10, 64)
		exitError(err)
		roomName := getArg(3)
		roomOptions := proto_lobby.RoomOptions{}
		err = json.Unmarshal([]byte(getArg(4)), &roomOptions)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.Uint64(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.CreateRoom(&auth, roomName, &roomOptions)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "join_room":
		userId, err := strconv.ParseUint(getArg(2), 10, 64)
		exitError(err)
		roomId := getArg(3)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.Uint64(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.JoinRoom(&auth, roomId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "leave_room":
		userId, err := strconv.ParseUint(getArg(2), 10, 64)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.Uint64(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.LeaveRoom(&auth)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "list_rooms":
		userId, err := strconv.ParseUint(getArg(2), 10, 64)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.Uint64(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.ListRooms(&auth)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "room_info":
		roomId := getArg(2)
		response, err := client.RoomInfo(roomId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	default:
		exitError(fmt.Errorf("Unknown action: %s", os.Args[1]))
	}
}

func exitError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getArg(index int) string {
	if len(os.Args) < index {
		return ""
	}
	return os.Args[index]
}
