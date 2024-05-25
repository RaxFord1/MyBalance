package personal

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	"MyBalance/internal/utils/secret"
	"fmt"
	"net/http"
)

type StatementResponse struct {
	Id              string `json:"id"`
	Time            int64  `json:"time"`
	Description     string `json:"description"`
	Mcc             int    `json:"mcc"`
	OriginalMcc     int    `json:"originalMcc"`
	Amount          int    `json:"amount"`
	OperationAmount int    `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int    `json:"commissionRate"`
	CashbackAmount  int    `json:"cashbackAmount"`
	Balance         int    `json:"balance"`
	Hold            bool   `json:"hold"`
	ReceiptId       string `json:"receiptId"`
}

func Statement(ctx context.Context, secretToken, account string, from, to int64) ([]StatementResponse, error) {
	url, err := ctx.GetString(projkeys.MonoApiUrl)
	if err != nil {
		return nil, err
	}

	finalUrl := fmt.Sprintf("%v/personal/statement/%v/%v/%v", url, account, from, to)

	request := &requesto.Request{
		Name:   "monoAPI-client-info",
		Url:    finalUrl,
		Method: http.MethodGet,
		Headers: map[string]string{
			"X-Token": secretToken,
		},
		MaskedHeaders: map[string]string{
			"X-Token": secret.ApplyMask(secretToken),
		},
	}
	result := &[]StatementResponse{}
	response := requesto.JsonResponse(result)

	if err = requesto.MakeRequest(ctx, request, response); err != nil {
		return nil, err
	}

	return *result, nil
}
