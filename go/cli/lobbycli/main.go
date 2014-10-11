package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"code.google.com/p/go.net/context"

	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_lobby"
	"github.com/opentarock/service-api/go/reqcontext"
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
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		response, err := client.CreateRoom(ctx, roomName, &roomOptions)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "join_room":
		userId := getArg(2)
		exitError(err)
		roomId := getArg(3)
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		response, err := client.JoinRoom(ctx, roomId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "leave_room":
		userId := getArg(2)
		exitError(err)
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		response, err := client.LeaveRoom(ctx)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "list_rooms":
		userId := getArg(2)
		exitError(err)
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		response, err := client.ListRooms(ctx)
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
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		response, err := client.StartGame(ctx)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "player_ready":
		userId := getArg(2)
		exitError(err)
		ctx = reqcontext.WithAuth(ctx, userId, "token")
		state := getArg(3)
		exitError(err)
		response, err := client.PlayerReady(ctx, state)
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
