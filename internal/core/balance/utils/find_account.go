package utils

import "MyBalance/pkg/services/api_monobank/personal"

func FindAccount(info *personal.ClientInfoStruct) personal.Account {
	for _, account := range info.Accounts {
		if account.CurrencyCode == 980 && account.Type == "black" {
			return account
		}
	}

	return info.Accounts[0]
}
