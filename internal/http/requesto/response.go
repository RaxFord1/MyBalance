package requesto

import (
	"context"
	"net/http"
)

type UnmarshalFunc func(ctx context.Context, data []byte, v interface{}) error

type ResponseResult struct {
	HttpCode int
	Header   http.Header
	Body     []byte
}

type Response struct {
	ResponseResult
	Response      interface{}
	UnmarshalFunc UnmarshalFunc
}

func JsonResponse(response interface{}) *Response {
	return &Response{
		Response:      response,
		UnmarshalFunc: ParseJSONResponse,
	}
}
