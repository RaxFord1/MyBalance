package mono_balance

import (
	"MyBalance/internal/utils/secret"
	"MyBalance/pkg/services/api_monobank/personal"
	"fmt"
	"strconv"
	"time"
)

func formatBalance(balance int) string {
	balanceStr := strconv.Itoa(balance)
	length := len(balanceStr)

	if length <= 2 {
		// If the balance is less than 100, prepend "0" or "00" as necessary
		switch length {
		case 1:
			return "0,0" + balanceStr
		case 2:
			return "0," + balanceStr
		}
	}

	// Insert comma before the last two digits
	formattedBalance := balanceStr[:length-2] + "," + balanceStr[length-2:]
	return formattedBalance
}

func formatCardInfo(account personal.Account) string {
	return fmt.Sprintf("%s\n%s\nОсталось на карте: %s\nБаланс: %s\nКредитный лимит: %s\n ",
		time.Now().Format("2006-01-02 15:04"),
		secret.ApplyMask(account.MaskedPan[0]),
		formatBalance(account.Balance-account.CreditLimit),
		formatBalance(account.Balance),
		formatBalance(account.CreditLimit),
	)
}
