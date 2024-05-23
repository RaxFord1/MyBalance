package requesto

import (
	"context"
	"net/http"
)

type MarshalFunc func(ctx context.Context, request interface{}) (body []byte, headers map[string]string, err error)

type Request struct {
	Name          string
	MarshalFunc   MarshalFunc
	Body          interface{}
	Method        string
	Headers       map[string]string
	MaskedHeaders map[string]string // todo: during logging, use masked headers
	Url           string
	DisableLog    bool

	Client *http.Client
}
