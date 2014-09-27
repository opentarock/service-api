package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/opentarock/service-api/go/client"
	"github.com/opentarock/service-api/go/proto_presence"
	"github.com/opentarock/service-api/go/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s ACTION [ARGS...]\n", os.Args[0])
		os.Exit(1)
	}
	client := client.NewPresenceClientNanomsg()
	err := client.Connect(service.MakeServiceAddress("localhost", service.PresenceServiceDefaultPort))
	exitError(err)
	defer client.Close()
	switch os.Args[1] {
	case "update_user_status":
		userId := getArg(2)
		statusString := getArg(3)
		status, ok := proto_presence.UpdateUserStatusRequest_Status_value[statusString]
		if !ok {
			log.Fatalf("Invalid status value: '%s'", statusString)
		}
		device := parseDevice(getArg(4))
		response, err := client.UpdateUserStatus(userId, proto_presence.UpdateUserStatusRequest_Status(status), device)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	case "get_user_devices":
		userId := getArg(2)
		response, err := client.GetUserDevices(userId)
		exitError(err)
		result, err := json.Marshal(response)
		exitError(err)
		fmt.Println(string(result))
	default:
		exitError(fmt.Errorf("Unknown action: %s", os.Args[1]))
	}
}

func parseDevice(deviceString string) *proto_presence.Device {
	parts := strings.SplitN(deviceString, ":", 2)
	if len(parts) != 2 {
		log.Fatalf("Invalid device string.")
	}
	device := &proto_presence.Device{}
	if parts[0] == "gcm" {
		device.Type = proto_presence.Device_ANDROID_GCM.Enum()
		device.GcmRegistrationId = &parts[1]
	} else {
		log.Fatalf("Unknown device type.")
	}
	return device
}

func exitError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getArg(index int) string {
	if len(os.Args) < index {
		return ""
	}
	return os.Args[index]
}
