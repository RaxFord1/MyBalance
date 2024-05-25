package utils

import (
	"MyBalance/internal/clients/monobank"
)

func FindAccount(info *monobank.ClientInfoStruct) monobank.Account {
	for _, account := range info.Accounts {
		if account.CurrencyCode == 980 && account.Type == "black" {
			return account
		}
	}

	return info.Accounts[0]
}
