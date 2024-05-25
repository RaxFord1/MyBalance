package tasks

import (
	"MyBalance/internal/core/balance/utils"
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/logger"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	"MyBalance/internal/telegram"
	"MyBalance/pkg/mono/mono_balance"
	"MyBalance/pkg/mono/mono_statement"
	"log"
	"time"
)

func getAndSendCurrentBalanceDaySummaryTask(ctx context.Context, b *telegram.Bot) error {
	f := func(ctx context.Context) {
		if err := getAndSendCurrentBalanceDaySummaryTask(ctx, b); err != nil {
			logger.PrintError(ctx, "getAndSendCurrentBalanceDaySummaryTask", err)
		}
	}

	go func() {
		defer logger.PanicAndCall(ctx, f, time.Hour)

		initialDelay := timeUntilNextDailyTarget(23, 55)
		log.Printf("Initial delay: %v\n\n", initialDelay)

		time.AfterFunc(initialDelay, func() {
			if err := GetAndSendCurrentBalanceDaySummary(ctx, b); err != nil {
				logger.PrintError(ctx, "UpdateLocationMerchant", err)
			}

			// send the message every 24 hours
			ticker := time.NewTicker(24 * time.Hour)
			for {
				<-ticker.C
				if err := GetAndSendCurrentBalanceDaySummary(ctx, b); err != nil {
					logger.PrintError(ctx, "UpdateLocationMerchant", err)
				}
			}
		})
	}()

	return nil
}

func GetAndSendCurrentBalanceDaySummary(ctx context.Context, b *telegram.Bot) error {
	if b == nil {
		return requesto.InternalError.NewWithMsg(ctx, "bot is nil")
	}
	users := db.GetUsers(ctx)
	now := time.Now()
	if len(users) == 0 {
		return nil
	}

	// TODO: now we have only 1 mono-api-key, but in future we can have multiple api keys, so this has to be modified
	ctx.SetString(projkeys.ClientID, users[0])
	balance, err := mono_balance.GetBalance(ctx)
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	state, err := mono_statement.GetStatement(ctx)
	if err != nil {
		return err
	}

	for _, v := range users {
		summary := utils.GenerateDaySummaryString(now, balance, state)
		if _, err = b.Send(ctx, v, summary); err != nil {
			return err
		}
	}

	return nil
}
