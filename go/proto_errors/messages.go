package proto_errors

import (
	"fmt"

	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	ErrorResponseMessage = 0x0F0101
)

func (m *ErrorResponse) GetMessageType() proto.Type {
	return ErrorResponseMessage
}

func NewMalformedMessageUnpack() *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_MALFORMED_MESSAGE.Enum(),
		Description: pbuf.String("Unable to unpack the message."),
	}
}

func NewInternalError(msg string) *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_INTERNAL_ERROR.Enum(),
		Description: &msg,
	}
}

func NewInternalErrorResponse() *ErrorResponse {
	return NewInternalError("There was a problem sending a response.")
}

func NewMissingHeader(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_MISSING_HEADER.Enum(),
		Description: pbuf.String(fmt.Sprintf("Missing header with type=%X", ty)),
	}
}

func NewUnsupportedMessage(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_INTERNAL_ERROR.Enum(),
		Description: pbuf.String(fmt.Sprintf("Message of type=%X is not supported by this service.", ty)),
	}
}

func NewMalformedMessage(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_MALFORMED_MESSAGE.Enum(),
		Description: pbuf.String(fmt.Sprintf("Message of type=%X could no not be decoded.", ty)),
	}
}

func NewEmptyMessage() *ErrorResponse {
	return &ErrorResponse{
		Error:       Error_EMPTY_MESSAGE.Enum(),
		Description: pbuf.String("Unexpected empty message"),
	}
}
