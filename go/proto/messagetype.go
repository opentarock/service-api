package proto

import pbuf "code.google.com/p/gogoprotobuf/proto"

type MessageType interface {
	GetMessageType() Type
}

type ProtobufMessage interface {
	MessageType
	pbuf.Message
}
