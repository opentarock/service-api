package service

import "fmt"

const (
	PresenceServiceDefaultPort = 9001
)

func MakeServiceBindAddress(port uint) string {
	return fmt.Sprintf("tcp://*:%d", port)
}

func MakeServiceAddress(host string, port uint) string {
	return fmt.Sprintf("tcp://%s:%d", host, port)
}
