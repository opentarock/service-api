package service

import (
	"io"

	"github.com/opentarock/service-api/go/proto"
)

type Service interface {
	Addhandler(messageType proto.Type, handler MessageHandler)
	Start() error
	io.Closer
}
