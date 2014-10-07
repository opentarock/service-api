package clientutil

import (
	"code.google.com/p/go.net/context"
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

func DoRequest(
	ctx context.Context,
	client *ReqClient,
	request proto.ProtobufMessage,
	response proto.ProtobufMessage,
	headers ...proto.ProtobufMessage) error {

	req, err := client.Request(request, headers...)
	defer req.Cancel()
	if err != nil {
		return err
	}

	select {
	case responseMsg := <-req.Done():
		if responseMsg == nil {
			return req.Err()
		} else {
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
	case <-ctx.Done():
		req.Cancel()
		return ctx.Err()
	}
}
