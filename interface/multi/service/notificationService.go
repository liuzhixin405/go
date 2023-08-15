package service

import (
	"context"
	"fmt"
)

type Notifier struct{}

func New() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyTopScore(ctx context.Context, id string, score int) error {
	fmt.Printf("Sending notification to ID %q about their high score of %d\n", id, score)
	return nil
}

func (n *Notifier) NotifyPasswordUpdate(ctx context.Context, id string) error {
	fmt.Printf("Sending notification to ID %q about their password being updated\n", id)
	return nil
}
