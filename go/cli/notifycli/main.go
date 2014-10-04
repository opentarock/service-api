package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"code.google.com/p/go.net/context"

	"github.com/docopt/docopt-go"
	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_notify"
	"github.com/opentarock/service-api/go/service"
)

const timeout = 5 * time.Second

func main() {
	usage := `Notify cli client.

Usage:
	notifycli message_users [options] <json> <user_ids>...

Options:
	-h --help      Show this screen.
	--host=<host>  Host to connect to. [default: 127.0.0.1]
	--port=<port>  Port to connect to. [default: 8001]`

	args, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		os.Exit(1)
	}
	log.SetOutput(ioutil.Discard)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client := client.NewNotifyClientNanomsg()
	port, err := strconv.ParseUint(args["--port"].(string), 10, 16)
	if err != nil {
		fmt.Print("Port number must be an unsigned integer")
		os.Exit(1)
	}
	err = client.Connect(service.MakeServiceAddress(args["--host"].(string), uint(port)))
	exitError(err)
	switch {
	case args["message_users"].(bool):
		userIds := args["<user_ids>"].([]string)
		var data string
		if j, ok := args["<json>"].(string); ok {
			data = j
		}
		var jsonMap map[string]interface{}
		err = json.Unmarshal([]byte(data), &jsonMap)
		if err != nil {
			displayError(err)
		}
		jsonMsg, _ := proto_notify.NewJsonMessage(jsonMap)
		_, err := client.MessageUsers(ctx, jsonMsg, userIds...)
		if err != nil {
			displayError(err)
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func displayError(err error) {
	fmt.Print("Problem sending message: ")
	exitError(err)
}

func exitError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
