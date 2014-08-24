package proto_oauth2

import (
	"code.google.com/p/gogoprotobuf/proto"

	"github.com/arjantop/oauth2-util"
)

func NewInvalidClientError(msg string) *ErrorResponse {
	return &ErrorResponse{
		Error:            proto.String(oauth2.ErrorInvalidClient),
		ErrorDescription: &msg,
	}
}
