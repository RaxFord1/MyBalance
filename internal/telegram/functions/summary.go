package functions

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/tasks"
	"MyBalance/internal/telegram"
	tele "gopkg.in/telebot.v3"
)

func Summary(ctx context.Context, c tele.Context) error {

	bot := telegram.BotFromTeleBot(ctx, c.Bot())

	err := tasks.GetAndSendCurrentBalanceDaySummary(ctx, bot)
	if err != nil {
		return c.Send(err.Error())
	}

	return nil
}
