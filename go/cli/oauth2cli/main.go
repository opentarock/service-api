package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_oauth2"
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
	client, err := client.NewOauth2ClientNanomsg()
	exitError(err)
	switch os.Args[1] {
	case "access_token":
		clientId := getArg(2)
		clientSecret := getArg(3)
		accessTokenRequest := proto_oauth2.AccessTokenRequest{}
		err = json.Unmarshal([]byte(getArg(4)), &accessTokenRequest)
		exitError(err)
		response, err := client.GetAccessToken(clientId, clientSecret, &accessTokenRequest)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	default:
		exitError(fmt.Errorf("Unknown action: %s", os.Args[1]))
	}
}
