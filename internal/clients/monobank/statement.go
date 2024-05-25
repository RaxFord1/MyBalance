package monobank

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
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

func (c *Client) Statement(ctx context.Context, account string, from, to int64) ([]StatementResponse, error) {
	finalUrl := fmt.Sprintf("%v/personal/statement/%v/%v/%v", c.baseURL, account, from, to)

	request := &requesto.Request{
		Name:   "monoAPI-client-info",
		Url:    finalUrl,
		Method: http.MethodGet,
		Headers: map[string]string{
			"X-Token": c.apiKey,
		},
		MaskedHeaders: map[string]string{
			"X-Token": secret.ApplyMask(c.apiKey),
		},
	}
	result := &[]StatementResponse{}
	response := requesto.JsonResponse(result)

	if err := requesto.MakeRequest(ctx, request, response); err != nil {
		return nil, err
	}

	return *result, nil
}
