package requesto

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func unmarshal(ctx context.Context, unmarshalFunc UnmarshalFunc, data []byte, v interface{}) error {
	if unmarshalFunc == nil {
		return ParseJSONResponse(ctx, data, v)
	}
	return unmarshalFunc(ctx, data, v)
}

func MakeRequest(ctx context.Context, request *Request, response *Response) error {
	req, err := http.NewRequestWithContext(ctx, request.Method, request.Url, nil)
	if err != nil {
		return FailedRequestCreation.New(ctx)
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	var client *http.Client
	if request.Client == nil {
		client = &http.Client{}
	} else {
		client = request.Client
	}

	resp, err := client.Do(req)
	if err != nil {
		// todo: maybe fill in response *Response (?)
		// and also could parse errors
		return err
	}
	defer resp.Body.Close()

	response.HttpCode = resp.StatusCode
	response.Header = resp.Header

	if response.HttpCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	response.Body, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err = unmarshal(ctx, response.UnmarshalFunc, response.Body, response.Response); err != nil {
		return err
	}

	return nil
}
