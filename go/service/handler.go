package service

import "github.com/opentarock/service-api/go/proto"

type MessageHandler interface {
	HandleMessage(msg *proto.Message) proto.CompositeMessage
}

type MessageHandlerFunc func(msg *proto.Message) proto.CompositeMessage

func (f MessageHandlerFunc) HandleMessage(msg *proto.Message) proto.CompositeMessage {
	return f(msg)
}
