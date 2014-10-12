package reqcontext

import (
	"time"

	"code.google.com/p/go.net/context"
	pbuf "code.google.com/p/gogoprotobuf/proto"
	"gopkg.in/inconshreveable/log15.v2"

	"github.com/opentarock/service-api/go/log"
	"github.com/opentarock/service-api/go/proto"
	"github.com/opentarock/service-api/go/proto_headers"
)

type key int

const (
	reqCorrKey key = iota
	authKey
)

func WithRequest(
	ctx context.Context,
	msg *proto.Message,
	defaultTimeout time.Duration) (context.Context, context.CancelFunc) {

	var corrHeader proto_headers.RequestCorrelationHeader
	found, err := msg.Header.Unmarshal(&corrHeader)
	if found && err == nil {
		ctx = context.WithValue(ctx, reqCorrKey, corrHeader)
	}

	var authHeader proto_headers.AuthorizationHeader
	found, err = msg.Header.Unmarshal(&authHeader)
	if found && err == nil {
		ctx = context.WithValue(ctx, authKey, authHeader)
	}

	var timeoutHeader proto_headers.TimeoutHeader
	found, err = msg.Header.Unmarshal(&timeoutHeader)
	var deadline time.Time
	if found && err == nil {
		deadline = timeoutHeader.DeadlineTime()
	} else {
		deadline = time.Now().Add(defaultTimeout)
	}
	return context.WithDeadline(ctx, deadline)
}

func WithCorrId(ctx context.Context, corrId proto_headers.RequestCorrelationHeader) context.Context {
	return context.WithValue(ctx, reqCorrKey, corrId)
}

func CorrIdFromContext(ctx context.Context) (*proto_headers.RequestCorrelationHeader, bool) {
	h, ok := ctx.Value(reqCorrKey).(proto_headers.RequestCorrelationHeader)
	return &h, ok
}

func ContextLogger(ctx context.Context, args ...interface{}) log15.Logger {
	corr, ok := CorrIdFromContext(ctx)
	if !ok {
		corr.CorrelationId = pbuf.String("none")
	}
	args = append(args, "corr_id")
	args = append(args, corr.GetCorrelationId())
	return log.New(args...)
}

func WithAuth(ctx context.Context, userId string, accessToken string) context.Context {
	auth := proto_headers.NewAutorizationHeader(userId, accessToken)
	return context.WithValue(ctx, authKey, *auth)
}

func AuthFromContext(ctx context.Context) (*proto_headers.AuthorizationHeader, bool) {
	h, ok := ctx.Value(authKey).(proto_headers.AuthorizationHeader)
	return &h, ok
}
