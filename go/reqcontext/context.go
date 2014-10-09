package reqcontext

import (
	"os"
	"time"

	"code.google.com/p/go.net/context"
	pbuf "code.google.com/p/gogoprotobuf/proto"
	log "gopkg.in/inconshreveable/log15.v2"

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

	var timeoutHeader proto_headers.TimeoutHeader
	found, err = msg.Header.Unmarshal(&timeoutHeader)
	var timeout time.Duration
	if found && err == nil {
		timeout = timeoutHeader.Duration()
	} else {
		timeout = defaultTimeout
	}
	return context.WithTimeout(ctx, timeout)
}

func CorrIdFromContext(ctx context.Context) (*proto_headers.RequestCorrelationHeader, bool) {
	h, ok := ctx.Value(reqCorrKey).(proto_headers.RequestCorrelationHeader)
	return &h, ok
}

func ContextLogger(ctx context.Context) log.Logger {
	corr, ok := CorrIdFromContext(ctx)
	if !ok {
		corr.CorrelationId = pbuf.String("none")
	}
	logger := log.New(log.Ctx{"corr_id": corr.GetCorrelationId()})
	logger.SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))
	return logger
}

func WithAuth(ctx context.Context, userId string, accessToken string) context.Context {
	auth := proto_headers.NewAutorizationHeader(userId, accessToken)
	return context.WithValue(ctx, authKey, auth)
}

func AuthFromContext(ctx context.Context) (*proto_headers.AuthorizationHeader, bool) {
	h, ok := ctx.Value(authKey).(proto_headers.AuthorizationHeader)
	return &h, ok
}
