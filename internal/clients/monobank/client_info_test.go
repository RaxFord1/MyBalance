package monobank_test

import (
	"MyBalance/internal/clients/monobank"
	"MyBalance/internal/core"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"log"
	"testing"
)

func TestClientInfo(t *testing.T) {
	ctx := context.Named("test-monoAPI-client-info")
	err := core.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	apiKey, err := ctx.GetString(projkeys.MonoApiKey)
	if err != nil {
		t.Fatal(err)
	}

	url, err := ctx.GetString(projkeys.MonoApiUrl)
	if err != nil {
		t.Fatal(err)
	}

	mbClient := monobank.NewClient(url, apiKey)

	got, err := mbClient.ClientInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(got)
}
