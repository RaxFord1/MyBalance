package mono_statement

import (
	"MyBalance/internal/core/balance/utils"
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"MyBalance/pkg/services/api_monobank/personal"
	"strings"
	"time"
)

func formatStatement(history []personal.StatementResponse) string {
	sb := strings.Builder{}

	for i := range history {
		sb.WriteString(history[i].Description)
		sb.WriteString("\n")
		sb.WriteString(utils.FormatBalance(history[i].Amount))
		sb.WriteString("\n")
		sb.WriteString(time.Unix(history[i].Time, 0).Format("2006-01-02 15:04:05"))
		sb.WriteString("\n")
		sb.WriteString("\n")
	}

	return sb.String()
}

func GetStatement(ctx context.Context) (string, error) {
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

	card, err := db.GetCard(ctx, clientId)
	if err != nil {
		return "", err
	}

	start, end := GetTimeStartAndNowUnix()

	info, err := personal.Statement(ctx, apiKey, card, start, end)
	if err != nil {
		return "", err
	}

	return formatStatement(info), nil
}
