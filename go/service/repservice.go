package service

import (
	"fmt"
	"log"

	nmsg "github.com/op/go-nanomsg"
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_errors"
)

type messageHandlers map[proto.Type]MessageHandler

type RepService struct {
	Address         string
	socket          *nmsg.RepSocket
	messageHandlers messageHandlers
	headerHandlers  messageHandlers
}

func NewRepService(bind string) *RepService {
	return &RepService{
		Address:         bind,
		messageHandlers: make(messageHandlers),
		headerHandlers:  make(messageHandlers),
	}
}

func (s *RepService) AddHandler(messageType proto.Type, handler MessageHandler) {
	log.Printf("Adding handler for message type: %X", uint64(messageType))
	s.messageHandlers[messageType] = handler
}

func (s *RepService) AddHeaderHandler(messageType proto.Type, handler MessageHandler) {
	log.Printf("Adding handler for header message type: %X", uint64(messageType))
	s.headerHandlers[messageType] = handler
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
				sendResponse(socket, proto_errors.NewEmptyMessage())
				continue
			}
			msg, err := proto.Parse(recvData)
			if err != nil {
				log.Printf("Error parsing message: %s", err)
				sendResponse(socket, proto_errors.NewMalformedMessageUnpack())
				continue
			}
			if handled, err := s.messageHandlers.handleType(socket, msg.Type, msg); handled {
				if err != nil {
					log.Println(err)
					sendResponse(socket, proto_errors.NewInternalErrorResponse())
					continue
				}
			} else {
				var handled bool
				for msgType, _ := range msg.Header {
					handled, err = s.headerHandlers.handleType(socket, msgType, msg)
					if err != nil {
						log.Println(err)
						sendResponse(socket, proto_errors.NewInternalErrorResponse())
						break
					} else if handled {
						break
					}
				}
				if !handled {
					log.Printf("Unknown message type: %s", msg.Type)
					sendResponse(socket, proto_errors.NewUnsupportedMessage(msg.Type))
				}
			}
		}
	}()
	return nil
}

func (h messageHandlers) handleType(socket *nmsg.RepSocket, msgType proto.Type, msg *proto.Message) (bool, error) {
	if handler, ok := h[msgType]; ok {
		response := handler.HandleMessage(msg)
		responseMsg, err := response.Marshal()
		if err != nil {
			return true, fmt.Errorf("Error marshalling response: %s", err)
		}
		responseData, err := responseMsg.Pack()
		if err != nil {
			return true, fmt.Errorf("Error packing response message: %s", err)
		}
		_, err = socket.Send(responseData, 0)
		if err != nil {
			return true, fmt.Errorf("Error sending response: %s", err)
		}
		return true, nil
	}
	return false, nil
}

func sendResponse(socket *nmsg.RepSocket, msg proto.ProtobufMessage) {
	msgResponse := proto.MarshalForce(msg)
	responseData, err := msgResponse.Pack()
	if err != nil {
		log.Panicf("Error packing response message: %s", err)
	}
	_, err = socket.Send(responseData, 0)
	if err != nil {
		log.Println("Error sending response: %s", err)
	}
}

func (s *RepService) Close() error {
	log.Printf("Closing service on: %s", s.Address)
	return s.socket.Close()
}
