package clientutil

import (
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_errors"
)

func TryDecodeError(msg *proto.Message, response proto.ProtobufMessage) error {
	if msg.Type != response.GetMessageType() {
		var responseError proto_errors.ErrorResponse
		err := msg.Unmarshal(response)
		if err != nil {
			return err
		}
		return &responseError
	}
	return nil
}
