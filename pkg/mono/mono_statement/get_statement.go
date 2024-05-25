package mono_statement

import (
	"MyBalance/internal/clients/monobank"
	"MyBalance/internal/core/balance/utils"
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"strings"
	"time"
)

func formatStatement(history []monobank.StatementResponse) string {
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

	url, err := ctx.GetString(projkeys.MonoApiUrl)
	if err != nil {
		return "", err
	}

	mbClient := monobank.NewClient(url, apiKey)

	info, err := mbClient.Statement(ctx, card, start, end)
	if err != nil {
		return "", err
	}

	return formatStatement(info), nil
}
