package functions

import (
	"MyBalance/internal/http/context"
	"MyBalance/pkg/mono/mono_balance"
	tele "gopkg.in/telebot.v3"
)

func Balance(ctx context.Context, c tele.Context) error {
	balance, err := mono_balance.GetBalance(ctx)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Send(balance)
}
