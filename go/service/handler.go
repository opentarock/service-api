package service

import "github.com/opentarock/service-api/go/proto"

type MessageHandler interface {
	HandleMessage(msg *proto.Message) *proto.Message
}

type MessageHandlerFunc func(msg *proto.Message) *proto.Message

func (f MessageHandlerFunc) HandleMessage(msg *proto.Message) *proto.Message {
	return f(msg)
}
