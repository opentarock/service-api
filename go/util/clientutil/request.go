package clientutil

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_errors"
	"github.com/opentarock/service-api/go/proto_headers"
	"github.com/opentarock/service-api/go/reqcontext"
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

	headers = headersFromContext(ctx, headers)

	req, err := client.Request(request, headers...)
	if err != nil {
		return err
	}
	defer req.Cancel()

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

func headersFromContext(ctx context.Context, headers []proto.ProtobufMessage) []proto.ProtobufMessage {
	deadline, ok := ctx.Deadline()
	if ok {
		headers = append(headers, proto_headers.NewDeadlineTimeoutHeader(deadline))
	}

	reqCorr, ok := reqcontext.CorrIdFromContext(ctx)
	if ok {
		headers = append(headers, reqCorr)
	}

	auth, ok := reqcontext.AuthFromContext(ctx)
	if ok {
		headers = append(headers, auth)
	}

	return headers
}
