package core

import (
	"MyBalance/internal/config"
	"MyBalance/internal/context"
	"MyBalance/internal/core/db"
	"MyBalance/internal/telegram"
	"MyBalance/pkg/mono"
)

func Init(ctx context.Context) error {

	if err := config.Init(ctx); err != nil {
		return err
	}

	if err := db.Init(ctx); err != nil {
		return err
	}

	if err := mono.Init(ctx); err != nil {
		return err
	}

	if err := telegram.Init(ctx); err != nil {
		return err
	}

	return nil
}
