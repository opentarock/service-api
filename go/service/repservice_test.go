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

func TestHandlerIsUsed(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9000")
	req.Connect("tcp://localhost:9000")
	called := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) *proto.Message {
			called = true
			return proto.NewMessage(1, []byte{})
		}))
	go func() {
		repService.Start()
	}()

	req.Send([]byte{1}, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, called, true)
	defer repService.Close()
}

func TestOnlyLastAddedHandlerForTypeIsUsed(t *testing.T) {
	req, err := nmsg.NewReqSocket()
	assert.Nil(t, err)
	repService := service.NewRepService("tcp://*:9000")
	req.Connect("tcp://localhost:9000")
	calledFirst := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) *proto.Message {
			calledFirst = true
			return proto.NewMessage(1, []byte{})
		}))
	calledSecond := false
	repService.AddHandler(1,
		service.MessageHandlerFunc(func(data *proto.Message) *proto.Message {
			calledSecond = true
			return proto.NewMessage(1, []byte{})
		}))
	go func() {
		repService.Start()
	}()

	req.Send([]byte{1}, 0)
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, calledFirst, false, "First handler should be owerwritten")
	assert.Equal(t, calledSecond, true)
	defer repService.Close()
}
