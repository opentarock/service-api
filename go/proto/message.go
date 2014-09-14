package proto

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"

	pbuf "code.google.com/p/gogoprotobuf/proto"
)

type Type uint64

type Header map[Type][]byte

func (h Header) Set(key Type, value []byte) {
	h[key] = value
}

func (h Header) Marshal(r ProtobufMessage) error {
	data, err := pbuf.Marshal(r)
	if err != nil {
		return ErrorMarshalling(r, err)
	}
	h.Set(r.GetMessageType(), data)
	return nil
}

func (h Header) Get(key Type) []byte {
	return h[key]
}

func (h Header) Unmarshal(r ProtobufMessage) (bool, error) {
	data := h.Get(r.GetMessageType())
	if data == nil {
		return false, ErrorMissingHeader(r)
	}
	err := pbuf.Unmarshal(data, r)
	if err != nil {
		return true, ErrorUnmarshalling(r, err)
	}
	return true, nil
}

type Message struct {
	Header Header
	Type   Type
	Data   []byte
}

func NewMessage(ty Type, data []byte) *Message {
	return &Message{
		Header: make(Header),
		Type:   ty,
		Data:   data,
	}
}

func Marshal(r ProtobufMessage) (*Message, error) {
	data, err := pbuf.Marshal(r)
	if err != nil {
		return nil, ErrorMarshalling(r, err)
	}
	return NewMessage(r.GetMessageType(), data), nil
}

func MarshalForce(m ProtobufMessage) *Message {
	msg, err := Marshal(m)
	if err != nil {
		log.Panicln(err)
	}
	return msg
}

func MarshalHeaders(r ProtobufMessage, h []ProtobufMessage) (*Message, error) {
	msg, err := Marshal(r)
	if err != nil {
		return nil, err
	}
	for _, header := range h {
		err = msg.Header.Marshal(header)
		if err != nil {
			return nil, err
		}
	}
	return msg, nil
}

func Parse(packetData []byte) (*Message, error) {
	buf := bytes.NewReader(packetData)
	pos := 0
	var message Message
	message.Header = make(Header)
	for pos < len(packetData) {
		length, err := binary.ReadUvarint(buf)
		if err != nil {
			return nil, err
		}
		ty, err := binary.ReadUvarint(buf)
		if err != nil {
			return nil, err
		}
		newPos := pos + int(length) + varintSize(length)
		// data is everything from the current position minus the length of encoded
		// part length and length of encoded type.
		data := packetData[pos+varintSize(length)+varintSize(ty) : newPos]
		_, err = buf.Seek(int64(newPos), 0)
		if err != nil {
			return nil, err
		}
		// if we read the last part its the actual message
		if len(packetData) == newPos {
			message.Type = Type(ty)
			message.Data = data
		} else {
			message.Header.Set(Type(ty), data)
		}
		pos = newPos
	}
	return &message, nil
}

func (p *Message) Size() int {
	size := 0
	for ty, data := range p.Header {
		size += chunkSize(ty, data)
	}
	size += chunkSize(p.Type, p.Data)
	return size
}

func chunkSize(ty Type, data []byte) int {
	size := partSize(ty, data)
	return varintSize(uint64(size)) + size
}

func partSize(ty Type, data []byte) int {
	typeSize := varintSize(uint64(ty))
	return typeSize + len(data)
}

func (p *Message) Pack() ([]byte, error) {
	w := bytes.NewBuffer(make([]byte, 0, p.Size()))
	for ty, data := range p.Header {
		packChunk(w, ty, data)
	}
	packChunk(w, p.Type, p.Data)
	return w.Bytes(), nil
}

func packChunk(w *bytes.Buffer, ty Type, data []byte) error {
	msgSize := partSize(ty, data)
	_, err := putUvarint(w, uint64(msgSize))
	if err != nil {
		return err
	}
	_, err = putUvarint(w, uint64(ty))
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// PutUvarint encodes a uint64 into ByteBuffer and returns the number of
// bytes written.
// Based on golang implementation (http://golang.org/pkg/encoding/binary/#PutUvarint).
func putUvarint(w io.ByteWriter, x uint64) (int, error) {
	i := 0
	for x >= 0x80 {
		err := w.WriteByte(byte(x) | 0x80)
		if err != nil {
			return i, err
		}
		x >>= 7
		i++
	}
	err := w.WriteByte(byte(x))
	if err != nil {
		return i, err
	}
	return i + 1, nil
}

func varintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func (m *Message) Unmarshal(r ProtobufMessage) error {
	if m.Type != r.GetMessageType() {
		return ErrorTypeMismatch(r, m.Type)
	}
	err := pbuf.Unmarshal(m.Data, r)
	if err != nil {
		return ErrorUnmarshalling(r, err)
	}
	return nil
}

type CompositeMessage struct {
	Message ProtobufMessage
	Headers []ProtobufMessage
}

func (m CompositeMessage) Marshal() (*Message, error) {
	return MarshalHeaders(m.Message, m.Headers)
}
