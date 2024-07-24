package notify

import (
	"context"
)

type Default struct {
}

func NewDefault() *Default {
	return &Default{}
}

func (d Default) Send(ctx context.Context, subject, message string) error {
	return nil
}
