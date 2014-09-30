package clientutil

import (
	"log"
	"time"

	nmsg "github.com/op/go-nanomsg"
	"github.com/opentarock/service-api/go/proto"
)

type ReqClient struct {
	socket *nmsg.ReqSocket
}

func NewReqClient() *ReqClient {
	socket, err := nmsg.NewReqSocket()
	if err != nil {
		log.Panicf("Error creating req socket: %s", err)
	}
	const timeout = 1 * time.Second
	err = socket.SetSendTimeout(timeout)
	err = socket.SetRecvTimeout(timeout)
	if err != nil {
		log.Panicf("Error setting socket timeout: %s", err)
	}
	return &ReqClient{
		socket: socket,
	}
}

func (s *ReqClient) Request(
	request proto.ProtobufMessage,
	headers ...proto.ProtobufMessage) (*proto.Message, error) {

	msg, err := proto.MarshalHeaders(request, headers)
	if err != nil {
		return nil, err
	}

	err = s.send(msg)
	if err != nil {
		return nil, err
	}
	responseMsg, err := s.recv()
	if err != nil {
		return nil, err
	}

	return responseMsg, nil
}

func (c *ReqClient) Connect(address string) error {
	endpoint, err := c.socket.Connect(address)
	if err != nil {
		return err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)
	return nil
}

func (s *ReqClient) send(msg *proto.Message) error {
	packed, err := msg.Pack()
	if err != nil {
		return err
	}
	_, err = s.socket.Send(packed, 0)
	return err
}

func (s *ReqClient) recv() (*proto.Message, error) {
	responseData, err := s.socket.Recv(0)
	if err != nil {
		return nil, err
	}
	responseMsg, err := proto.Parse(responseData)
	if err != nil {
		return nil, err
	}
	return responseMsg, err
}

func (s *ReqClient) Close() error {
	return s.socket.Close()
}
