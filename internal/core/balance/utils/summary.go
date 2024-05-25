package utils

import (
	"strings"
	"time"
)

func GenerateDaySummaryString(tm time.Time, balance string, statement string) string {
	sb := strings.Builder{}
	sb.WriteString(tm.Format("2006-01-02"))
	sb.WriteString("\n")
	sb.WriteString("Остаток: ")
	sb.WriteString(balance)
	sb.WriteString("грн.\n")

	if statement == "" {
		sb.WriteString("Нет операций за этот день")
		return sb.String()
	}

	sb.WriteString(statement)

	return sb.String()
}
