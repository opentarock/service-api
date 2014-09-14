package proto

import (
	"fmt"
	"reflect"
)

func ErrorUnmarshalling(r ProtobufMessage, err error) error {
	return fmt.Errorf("Error unmarshalling %s: %s", reflect.TypeOf(r), err)
}

func ErrorMarshalling(r ProtobufMessage, err error) error {
	return fmt.Errorf("Error marshalling %s: %s", reflect.TypeOf(r), err)
}

func ErrorTypeMismatch(r ProtobufMessage, actual Type) error {
	return fmt.Errorf("Expecting %s (type=%X) but got message with type %d",
		reflect.TypeOf(r), r.GetMessageType(), actual)
}

func ErrorMissingHeader(r ProtobufMessage) error {
	return fmt.Errorf("Missing header %s (type=%X)", reflect.TypeOf(r), r.GetMessageType())
}