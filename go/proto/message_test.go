package proto_test

import (
	"testing"

	"github.com/opentarock/service-api/go/proto"
	"github.com/stretchr/testify/assert"
)

func makeSlice(n int, val byte) []byte {
	s := make([]byte, n)
	for i, _ := range s {
		s[i] = val
	}
	return s
}

func TestMessageIsEncoded(t *testing.T) {
	message := proto.NewMessage(128, makeSlice(127, 3))
	expected := makeSlice(131, 3)
	// message length
	expected[0] = 0x81
	expected[1] = 0x01
	// message type
	expected[2] = 0x80
	expected[3] = 0x01
	result, err := message.Pack()
	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)
}

func TestMessageWithHeadersIsEncoded(t *testing.T) {
	message := proto.NewMessage(1, makeSlice(1, 3))
	message.Header.Set(2, makeSlice(2, 4))
	message.Header.Set(3, makeSlice(3, 5))
	expected := makeSlice(12, 3)
	expected[0] = 0x03 // header 1 length
	expected[1] = 0x02 // header 1 type
	expected[2] = 0x04
	expected[3] = 0x04
	expected[4] = 0x04 // header 2 length
	expected[5] = 0x03 // header 2 type
	expected[6] = 0x05
	expected[7] = 0x05
	expected[8] = 0x05
	expected[9] = 0x02  // main message length
	expected[10] = 0x01 // main message type
	expected[11] = 0x03
	result, err := message.Pack()
	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)
}

func TestMessageEncodingAndDecodingIsIdentity(t *testing.T) {
	message := proto.NewMessage(1, makeSlice(1, 3))
	message.Header.Set(2, makeSlice(2, 4))
	message.Header.Set(3, makeSlice(3, 5))
	packed, err := message.Pack()
	assert.Nil(t, err)
	messageParsed, err := proto.Parse(packed)
	assert.Nil(t, err)
	assert.Equal(t, message, messageParsed)
}

func BenchmarkMessageEncode(b *testing.B) {
	message := proto.NewMessage(1, makeSlice(1, 3))
	message.Header.Set(2, makeSlice(2, 4))
	message.Header.Set(3, makeSlice(3, 5))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		message.Pack()
	}
}

func BenchmarkMessageDecode(b *testing.B) {
	message := proto.NewMessage(1, makeSlice(1, 3))
	message.Header.Set(2, makeSlice(2, 4))
	message.Header.Set(3, makeSlice(3, 5))
	packed, _ := message.Pack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Parse(packed)
	}
}
