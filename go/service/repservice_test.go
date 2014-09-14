package service_test

import (
	"testing"
	"time"

	_ "github.com/lib/pq"
	nmsg "github.com/op/go-nanomsg"
	"github.com/stretchr/testify/assert"

	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/service"
)

func TestMessageHandler(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9000")
	defer repService.Close()
	req.Connect("tcp://localhost:9000")
	called := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			called = true
			return proto.CompositeMessage{}
		}))
	go func() {
		repService.Start()
	}()

	data, err := proto.NewMessage(1, []byte{}).Pack()
	assert.Nil(t, err)
	req.Send(data, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, called, true)
}

func TestOnlyLastAddedHandlerForTypeIsUsed(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9001")
	defer repService.Close()
	req.Connect("tcp://localhost:9001")
	calledFirst := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledFirst = true
			return proto.CompositeMessage{}
		}))
	calledSecond := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledSecond = true
			return proto.CompositeMessage{}
		}))
	go func() {
		repService.Start()
	}()

	data, err := proto.NewMessage(1, []byte{}).Pack()
	assert.Nil(t, err)
	req.Send(data, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, calledFirst, false, "First handler should be owerwritten")
	assert.Equal(t, calledSecond, true)
}

func TestHeaderHandler(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9003")
	defer repService.Close()
	req.Connect("tcp://localhost:9003")
	calledFirst := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledFirst = true
			return proto.CompositeMessage{}
		}))
	calledSecond := false
	repService.AddHeaderHandler(2,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledSecond = true
			return proto.CompositeMessage{}
		}))
	go func() {
		repService.Start()
	}()

	msg := proto.NewMessage(3, []byte{})
	msg.Header.Set(2, []byte{})
	data, err := msg.Pack()
	assert.Nil(t, err)
	req.Send(data, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, calledFirst, false, "Message handler should not be executed")
	assert.Equal(t, calledSecond, true, "Header handler should be executed")
}

func TestMessageHandlerHasHigherPriority(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9004")
	defer repService.Close()
	req.Connect("tcp://localhost:9004")
	calledFirst := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledFirst = true
			return proto.CompositeMessage{}
		}))
	calledSecond := false
	repService.AddHeaderHandler(2,
		service.MessageHandlerFunc(func(data *proto.Message) proto.CompositeMessage {
			calledSecond = true
			return proto.CompositeMessage{}
		}))
	go func() {
		repService.Start()
	}()

	msg := proto.NewMessage(1, []byte{})
	msg.Header.Set(2, []byte{})
	data, err := msg.Pack()
	assert.Nil(t, err)
	req.Send(data, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, calledFirst, true, "Message handler has higher priority and should be called")
	assert.Equal(t, calledSecond, false, "Header handler is tryed only if there are no handlers for the message")
}
