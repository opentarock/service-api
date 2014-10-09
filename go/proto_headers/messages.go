package proto_headers

import (
	"math"
	"time"

	"code.google.com/p/go-uuid/uuid"
	pbuf "code.google.com/p/gogoprotobuf/proto"

	"github.com/opentarock/service-api/go/proto"
)

const (
	AuthorizationHeaderMessage      = 0x060101
	RequestCorrelationHeaderMessage = 0x060201
	TimeoutHeaderType               = 0x060301
)

func (m *AuthorizationHeader) GetMessageType() proto.Type {
	return AuthorizationHeaderMessage
}

func NewAutorizationHeader(userId string, accessToken string) *AuthorizationHeader {
	return &AuthorizationHeader{
		UserId:      &userId,
		AccessToken: &accessToken,
	}
}

func (h *RequestCorrelationHeader) GetMessageType() proto.Type {
	return RequestCorrelationHeaderMessage
}

func NewRequestCorrelationHeader() *RequestCorrelationHeader {
	return &RequestCorrelationHeader{
		CorrelationId: pbuf.String(uuid.New()),
	}
}

func (h *TimeoutHeader) GetMessageType() proto.Type {
	return TimeoutHeaderType
}

func (h *TimeoutHeader) Duration() time.Duration {
	if h.Timeout != nil {
		return time.Duration(h.GetTimeout()) * time.Nanosecond
	} else {
		return unixNanosec(h.GetDeadline()).Sub(time.Now())
	}
}

func unixNanosec(un uint64) time.Time {
	p := int64(math.Pow10(9))
	msec := int64(un) / p
	nsec := int64(un) % p
	return time.Unix(msec, nsec)
}

func NewTimeoutHeader(d time.Duration) *TimeoutHeader {
	return &TimeoutHeader{
		Timeout: pbuf.Uint64(uint64(d.Nanoseconds())),
	}
}

func NewDeadlineTimeoutHeader(t time.Time) *TimeoutHeader {
	return &TimeoutHeader{
		Deadline: pbuf.Uint64(uint64(t.UnixNano())),
	}
}
