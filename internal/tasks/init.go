package tasks

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/telegram"
)

var b *telegram.Bot

func Init(ctxOld context.Context) (err error) {
	ctx := context.NewFromPrev(ctxOld, "summary-task")

	b, err = telegram.NewBot(ctx)
	if err != nil {
		return err
	}

	if err = getAndSendCurrentBalanceDaySummaryTask(ctx, b); err != nil {
		return err
	}

	return
}
