package queue

import "context"

type Message struct {
	ID   string
	Body []byte
}

type Queue interface {
	Receive(ctx context.Context) (*Message, error)
	Delete(ctx context.Context, msg *Message) error
}
