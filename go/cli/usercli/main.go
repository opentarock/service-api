package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_user"
)

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

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s ACTION [ARGS...]\n", os.Args[0])
		os.Exit(1)
	}
	client, err := client.NewUserClientNanomsg()
	exitError(err)
	switch os.Args[1] {
	case "register_user":
		user := proto_user.User{}
		err = json.Unmarshal([]byte(getArg(2)), &user)
		exitError(err)
		redirectUri := getArg(3)
		registerResponse, err := client.RegisterUser(&user, redirectUri)
		exitError(err)
		result, err := json.Marshal(registerResponse)
		exitError(err)
		fmt.Println(string(result))
	case "auth_user":
		email := getArg(2)
		password := getArg(3)
		authResponse, err := client.AuthenticateUser(email, password)
		exitError(err)
		result, err := json.Marshal(authResponse)
		exitError(err)
		fmt.Println(string(result))
	default:
		exitError(fmt.Errorf("Unknown action: %s", os.Args[1]))
	}
}
