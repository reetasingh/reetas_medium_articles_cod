package main

import (
	"context"
	"fmt"
)

type RetryFunc func() error

type IReader interface {
	Read() error
}

type Reader struct {
}

func (r Reader) Read() error {
	return fmt.Errorf("error in read")
}

func Retry(ctx context.Context, f RetryFunc, retryAttempts int) error {
	var err error
	for i := 0; i < retryAttempts; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		err := f()
		if err == nil {
			return nil
		}
	}
	return err
}

func main() {
	reader := Reader{}
	Retry(context.Background(), reader.Read, 3)
}
