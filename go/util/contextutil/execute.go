package contextutil

import "code.google.com/p/go.net/context"

func Do(ctx context.Context, f func() error) error {
	c := make(chan error, 1)
	go func() { c <- f() }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
