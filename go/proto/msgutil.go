package proto

import "fmt"

func ErrorUnmarshalling(name string, err error) error {
	return fmt.Errorf("Error unmarshalling %s: err", name, err)
}

func ErrorTypeMismatch(name string, expected, actual Type) error {
	return fmt.Errorf("Expecting %s (type=%d) but got message with type %d", name, expected, actual)
}
