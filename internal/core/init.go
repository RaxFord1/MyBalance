package core

import (
	"MyBalance/internal/config"
	"MyBalance/internal/context"
	"MyBalance/internal/telegram"
	"MyBalance/pkg/mono_balance"
)

func Init(ctx context.Context) error {

	if err := config.Init(ctx); err != nil {
		return err
	}

	if err := mono_balance.Init(ctx); err != nil {
		return err
	}

	if err := telegram.Init(ctx); err != nil {
		return err
	}

	return nil
}
