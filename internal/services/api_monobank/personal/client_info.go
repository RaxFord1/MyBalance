package personal

import (
	"MyBalance/internal/context"
	"MyBalance/internal/http/requesto"
	"net/http"
)

type Account struct {
	Id           string   `json:"id"`
	SendId       string   `json:"sendId"`
	CurrencyCode int      `json:"currencyCode"`
	CashbackType string   `json:"cashbackType,omitempty"`
	Balance      int      `json:"balance"`
	CreditLimit  int      `json:"creditLimit"`
	MaskedPan    []string `json:"maskedPan"`
	Type         string   `json:"type"`
	Iban         string   `json:"iban"`
}

type Accounts []Account

type ClientInfoStruct struct {
	ClientId    string   `json:"clientId"`
	Name        string   `json:"name"`
	WebHookUrl  string   `json:"webHookUrl"`
	Permissions string   `json:"permissions"`
	Accounts    Accounts `json:"accounts"`
}

func ClientInfo(ctx context.Context, secretToken string) (*ClientInfoStruct, error) {
	request := &requesto.Request{
		Name:   "monoAPI-client-info",
		Url:    "https://api.monobank.ua/personal/client-info",
		Method: http.MethodGet,
		Headers: map[string]string{
			"X-Token": secretToken,
		},
		DisableLog: false,
	}
	responseXml := &ClientInfoStruct{}
	response := requesto.JsonResponse(responseXml)

	err := requesto.MakeRequest(ctx, request, response)
	if err != nil {
		return nil, err
	}

	return responseXml, nil
}
