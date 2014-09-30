package client

import "io"

type Client interface {
	Connect(address string) error
	io.Closer
}
