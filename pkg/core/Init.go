package core

import (
	"MyBalance/internal/context"
	"MyBalance/pkg/mono_balance"
)

func Init(ctx context.Context) error {

	if err := mono_balance.Init(ctx); err != nil {
		return err
	}

	return nil
}
