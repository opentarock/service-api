package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"code.google.com/p/go.net/context"
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
	client := client.NewLobbyClientNanomsg()
	err := client.Connect("tcp://localhost:7001")
	exitError(err)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer client.Close()
	switch os.Args[1] {
	case "create_room":
		userId := getArg(2)
		exitError(err)
		roomName := getArg(3)
		roomOptions := proto_lobby.RoomOptions{}
		err = json.Unmarshal([]byte(getArg(4)), &roomOptions)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.CreateRoom(ctx, &auth, roomName, &roomOptions)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "join_room":
		userId := getArg(2)
		exitError(err)
		roomId := getArg(3)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.JoinRoom(ctx, &auth, roomId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "leave_room":
		userId := getArg(2)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.LeaveRoom(ctx, &auth)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "list_rooms":
		userId := getArg(2)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.ListRooms(ctx, &auth)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "room_info":
		roomId := getArg(2)
		response, err := client.RoomInfo(ctx, roomId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "start_game":
		userId := getArg(2)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		response, err := client.StartGame(ctx, &auth)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "player_ready":
		userId := getArg(2)
		exitError(err)
		auth := proto_headers.AuthorizationHeader{
			UserId:      pbuf.String(userId),
			AccessToken: pbuf.String("token"),
		}
		state := getArg(3)
		exitError(err)
		response, err := client.PlayerReady(ctx, &auth, state)
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
