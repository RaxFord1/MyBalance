package mono_balance

import (
	"MyBalance/internal/clients/monobank"
	"MyBalance/internal/core/balance/utils"
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"MyBalance/internal/utils/secret"
	"fmt"
	"time"
)

type Balance struct {
	Credit  int
	Balance int
}

func formatCardInfo(account monobank.Account) string {
	return fmt.Sprintf("%s\n%s\nОсталось на карте: %s\nБаланс: %s\nКредитный лимит: %s\n ",
		time.Now().Format("2006-01-02 15:04"),
		secret.ApplyMask(account.MaskedPan[0]),
		utils.FormatBalance(account.Balance-account.CreditLimit),
		utils.FormatBalance(account.Balance),
		utils.FormatBalance(account.CreditLimit),
	)
}

func GetBalance(ctx context.Context) (string, error) {
	if err := LimitCheck(ctx, "mono-api-balance"); err != nil {
		return "", err
	}

	apiKey, err := ctx.GetString(projkeys.MonoApiKey)
	if err != nil {
		return "", err
	}

	clientId, err := ctx.GetString(projkeys.ClientID)
	if err != nil {
		return "", err
	}

	url, err := ctx.GetString(projkeys.MonoApiUrl)
	if err != nil {
		return "", err
	}

	mbClient := monobank.NewClient(url, apiKey)

	info, err := mbClient.ClientInfo(ctx)
	if err != nil {
		return "", err
	}

	account := utils.FindAccount(info)

	db.SetCard(ctx, clientId, account.Id)

	return formatCardInfo(account), nil
}
