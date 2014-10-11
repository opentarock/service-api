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
		ErrorCode:   ErrorCode_MALFORMED_MESSAGE.Enum(),
		Description: pbuf.String("Unable to unpack the message."),
	}
}

func NewInternalError(msg string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_INTERNAL_ERROR.Enum(),
		Description: &msg,
	}
}

func NewInternalErrorResponse() *ErrorResponse {
	return NewInternalError("There was a problem sending a response.")
}

func NewInternalErrorUnknown() *ErrorResponse {
	return NewInternalError("Unknown error occured.")
}

func NewMissingHeader(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_MISSING_HEADER.Enum(),
		Description: pbuf.String(fmt.Sprintf("Missing header with type=%X", ty)),
	}
}

func NewMissingFieldError(name string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_MISSING_FIELD.Enum(),
		Description: pbuf.String(fmt.Sprintf("Missing required field: '%s'", name)),
	}
}

func NewUnsupportedMessage(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_UNSUPPORTED_MESSAGE.Enum(),
		Description: pbuf.String(fmt.Sprintf("Message of type=%s is not supported by this service.", ty)),
	}
}

func NewMalformedMessage(ty proto.Type) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_MALFORMED_MESSAGE.Enum(),
		Description: pbuf.String(fmt.Sprintf("Message of type=%s could no not be decoded.", ty)),
	}
}

func NewEmptyMessage() *ErrorResponse {
	return &ErrorResponse{
		ErrorCode:   ErrorCode_EMPTY_MESSAGE.Enum(),
		Description: pbuf.String("Unexpected empty message"),
	}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s (type %s)", e.GetDescription(), e.GetErrorCode().String())
}
