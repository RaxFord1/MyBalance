package Mono

import (
	"MyBalance/internal/context"
	"MyBalance/pkg/services/api_monobank/personal"
	"os"
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
	info, err := personal.ClientInfo(ctx, os.Getenv("mono_api"))
	if err != nil {
		return "", err
	}

	account := findAccount(info)

	return formatCardInfo(account), nil
}
