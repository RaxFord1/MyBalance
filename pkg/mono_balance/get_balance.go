package mono_balance

import (
	"MyBalance/internal/context"
	"MyBalance/internal/projkeys"
	"MyBalance/pkg/services/api_monobank/personal"
)

type Balance struct {
	Credit  int
	Balance int
}

func findAccount(info *personal.ClientInfoStruct) personal.Account {
	for _, account := range info.Accounts {
		if account.CurrencyCode == 980 && account.Type == "black" {
			return account
		}
	}

	return info.Accounts[0]
}

func GetBalance(ctx context.Context) (string, error) {
	if err := LimitCheck(ctx, "mono-api-balance"); err != nil {
		return "", err
	}

	apiKey, err := ctx.GetString(projkeys.MonoApiKey)
	if err != nil {
		return "", err
	}

	info, err := personal.ClientInfo(ctx, apiKey)
	if err != nil {
		return "", err
	}

	account := findAccount(info)

	return formatCardInfo(account), nil
}
