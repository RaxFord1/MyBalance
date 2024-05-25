package core

import (
	"MyBalance/internal/config"
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
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
