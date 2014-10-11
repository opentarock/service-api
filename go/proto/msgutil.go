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
	return fmt.Errorf("Expecting %s (type=%s) but got message with type=%s",
		reflect.TypeOf(r), r.GetMessageType(), actual)
}

func ErrorMissingHeader(r ProtobufMessage) error {
	return fmt.Errorf("Missing header %s (type=%s)", reflect.TypeOf(r), r.GetMessageType())
}
