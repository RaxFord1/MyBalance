package monobank

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/utils/secret"
	"fmt"
	"log"
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

func (c *Client) ClientInfo(ctx context.Context) (*ClientInfoStruct, error) {
	if val, ok := ctx.Get("use_cache"); ok {
		if val == "true" {
			log.Println("using cache")
			return getCache(), nil
		}
	}

	finalUrl := fmt.Sprintf("%v/personal/client-info", c.baseURL)

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
	result := &ClientInfoStruct{}
	response := requesto.JsonResponse(result)

	if err := requesto.MakeRequest(ctx, request, response); err != nil {
		return nil, err
	}

	return result, nil
}
