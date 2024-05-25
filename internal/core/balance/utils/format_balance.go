package utils

import "strconv"

func FormatBalance(balance int) string {
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
