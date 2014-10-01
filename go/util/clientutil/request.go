package clientutil

import (
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_errors"
)

func TryDecodeError(msg *proto.Message, response proto.ProtobufMessage) error {
	if msg.Type != response.GetMessageType() {
		var responseError proto_errors.ErrorResponse
		err := msg.Unmarshal(&responseError)
		if err != nil {
			return err
		}
		return &responseError
	}
	return nil
}

func DoRequest(client *ReqClient, request proto.ProtobufMessage, response proto.ProtobufMessage) error {
	responseMsg, err := client.Request(request)
	if err != nil {
		return err
	}

	err = TryDecodeError(responseMsg, response)
	if err != nil {
		return err
	}

	err = responseMsg.Unmarshal(response)
	if err != nil {
		return err
	}

	return nil
}
