package mono

import (
	"MyBalance/internal/http/context"
	"MyBalance/pkg/mono/mono_balance"
	"MyBalance/pkg/mono/mono_statement"
)

func Init(ctx context.Context) error {
	if err := mono_balance.Init(ctx); err != nil {
		return err
	}

	if err := mono_statement.Init(ctx); err != nil {
		return err
	}

	return nil
}
