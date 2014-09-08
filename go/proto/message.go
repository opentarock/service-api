package proto

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Type uint64

type Header map[Type][]byte

func (h Header) Set(key Type, value []byte) {
	h[key] = value
}

func (h Header) Get(key Type) ([]byte, bool) {
	data, ok := h[key]
	return data, ok
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
