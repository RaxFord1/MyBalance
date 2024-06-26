package functions

import (
	"MyBalance/internal/http/context"
	"MyBalance/pkg/mono/mono_statement"
	tele "gopkg.in/telebot.v3"
)

func Statement(ctx context.Context, c tele.Context) error {
	balance, err := mono_statement.GetForTodayStatementString(ctx)
	if err != nil {
		return c.Send(err.Error())
	}

	if balance == "" {
		return c.Send("Нет истории транзакций за сегодня.")
	}

	return c.Send(balance)
}
