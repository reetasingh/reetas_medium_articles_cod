package main

import (
	"context"
	"fmt"
)

type IReader interface {
	Read() error
}

type Reader struct {
}

func (r Reader) Read() error {
	return fmt.Errorf("error in read")
}

type ReaderRetryDecorator struct {
	retryAttempts int
	reader        IReader
}

func (d ReaderRetryDecorator) Read() error {
	var err error
	for i := 0; i < d.retryAttempts; i++ {
		err := d.reader.Read()
		if err == nil {
			return nil
		}
	}
	return err
}

func (d ReaderRetryDecorator) ReadWithRetry(ctx context.Context, retryAttempts int) error {
	var err error
	for i := 0; i < retryAttempts; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		err := d.reader.Read()
		if err == nil {
			return nil
		}
	}
	return err
}

func main() {
	//retryReader := ReaderRetryDecorator{Reader{}}
	//retryReader.ReadWithRetry(context.Background(), 3)
	retryReader := ReaderRetryDecorator{3, Reader{}}
	retryReader.Read()
}
