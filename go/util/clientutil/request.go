package clientutil

import (
	"time"

	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_errors"
	"github.com/opentarock/service-api/go/proto_headers"
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

	deadline, ok := ctx.Deadline()
	if ok {
		headers = append(headers, proto_headers.NewTimeoutHeader(fromDeadline(deadline)))
	}

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

func fromDeadline(t time.Time) time.Duration {
	d := t.Sub(time.Now())
	if d < 0 {
		return 0
	}
	return d
}
