package reqcontext

import (
	"code.google.com/p/go.net/context"
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_headers"
)

type key int

const reqCorrKey key = 0

func NewContext(ctx context.Context, msg *proto.Message, excludeHeaders ...proto.Type) context.Context {
	var corrHeader proto_headers.RequestCorrelationHeader
	found, err := msg.Header.Unmarshal(&corrHeader)
	if !found || err != nil {
		return ctx
	}
	return context.WithValue(ctx, reqCorrKey, corrHeader)
}

func CorrIdFromContext(ctx context.Context) (*proto_headers.RequestCorrelationHeader, bool) {
	h, ok := ctx.Value(reqCorrKey).(proto_headers.RequestCorrelationHeader)
	return &h, ok
}
