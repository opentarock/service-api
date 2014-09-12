package service

import (
	"fmt"
	"log"

	nmsg "github.com/op/go-nanomsg"
	"github.com/opentarock/service-api/go/proto"
)

type RepService struct {
	Address         string
	socket          *nmsg.RepSocket
	messageHandlers map[proto.Type]MessageHandler
}

func NewRepService(bind string) *RepService {
	return &RepService{
		Address:         bind,
		messageHandlers: make(map[proto.Type]MessageHandler),
	}
}

func (s *RepService) AddHandler(messageType proto.Type, handler MessageHandler) {
	log.Printf("Adding handler for message type: %X", messageType)
	s.messageHandlers[messageType] = handler
}

func (s *RepService) Start() error {
	socket, err := nmsg.NewRepSocket()
	if err != nil {
		return fmt.Errorf("Error creating response socket: %s", err)
	}
	s.socket = socket

	endpoint, err := socket.Bind(s.Address)
	if err != nil {
		return fmt.Errorf("Error binding socket: %s", err)
	}
	log.Printf("Bound to endpoint: %s", endpoint.Address)

	go func() {
		for {
			recvData, err := socket.Recv(0)
			if err != nil {
				log.Printf("Error receiving message: %s", err)
				continue
			}
			if len(recvData) < 1 {
				log.Print("Unexpected empty message")
				continue
			}
			msg, err := proto.Parse(recvData)
			if err != nil {
				log.Printf("Error parsing message: %s", err)
				continue
			}
			if handler, ok := s.messageHandlers[msg.Type]; ok {
				response := handler.HandleMessage(msg)
				responseMsg, err := response.Marshal()
				if err != nil {
					log.Printf("Error marshalling response: %s", err)
					continue
				}
				responseData, err := responseMsg.Pack()
				if err != nil {
					log.Printf("Error packing response message: %s", err)
					continue
				}
				socket.Send(responseData, 0)
			} else {
				log.Printf("Unknown message type: %d", recvData[0])
			}
		}
	}()
	return nil
}

func (s *RepService) Close() error {
	log.Printf("Closing service on: %s", s.Address)
	return s.socket.Close()
}
