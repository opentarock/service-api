package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/service"
)

func main() {
	usage := `Google Cloud Messaging cli client.

Usage:
	gcmcli send_message [options] <registration_id>... [--data=<json>]

Options:
	-h --help      Show this screen.
	--host=<host>  Host to connect to. [default: "127.0.0.1"]
	--port=<port>  Port to connect to. [default: 11001]`

	args, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		fmt.Println("Error parsing arguments: ", err)
		os.Exit(1)
	}
	log.SetOutput(ioutil.Discard)
	client := client.NewGcmClientNanomsg()
	port, err := strconv.ParseUint(args["--port"].(string), 10, 16)
	if err != nil {
		fmt.Print("Port number must be an unsigned integer")
		os.Exit(1)
	}
	err = client.Connect(service.MakeServiceAddress(args["--host"].(string), uint(port)))
	exitError(err)
	switch {
	case args["send_message"].(bool):
		registrationIds := args["<registration_id>"].([]string)
		var data string
		if json, ok := args["--data"].(string); ok {
			data = json
		}
		response, err := client.SendMessage(registrationIds, data, nil)
		if err != nil {
			displayError(err)
		} else if response.ErrorCode != nil {
			displayError(errors.New(response.GetErrorCode().String()))
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
