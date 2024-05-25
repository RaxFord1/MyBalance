package utils

import (
	"strings"
)

func GenerateDaySummaryString(balance string, statement string) string {
	sb := strings.Builder{}

	sb.WriteString(balance)
	sb.WriteString("\n")

	if statement == "" {
		sb.WriteString("Нет операций за этот день")
		return sb.String()
	}

	sb.WriteString(statement)

	return sb.String()
}
