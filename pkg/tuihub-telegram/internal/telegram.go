package internal

import (
	"context"

	"github.com/nikoksr/notify/service/telegram"
)

func newClient(token string, chatID int64) (*telegram.Telegram, error) {
	client, err := telegram.New(token)
	if err != nil {
		return nil, err
	}
	client.AddReceivers(chatID)
	return client, nil
}

func Send(ctx context.Context, token string, chatID int64, subject, message string) error {
	client, err := newClient(token, chatID)
	if err != nil {
		return err
	}
	err = client.Send(ctx, subject, message)
	if err != nil {
		return err
	}
	return nil
}

func SendBatch(ctx context.Context, token string, chatID int64, messages map[string]string) error {
	client, err := newClient(token, chatID)
	if err != nil {
		return err
	}
	for subject, message := range messages {
		err = client.Send(ctx, subject, message)
		if err != nil {
			return err
		}
	}
	return nil
}
