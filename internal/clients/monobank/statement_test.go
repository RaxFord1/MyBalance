package monobank_test

import (
	"MyBalance/internal/clients/monobank"
	"MyBalance/internal/core"
	"MyBalance/internal/http/context"
	"MyBalance/internal/projkeys"
	"testing"
)

func TestStatement(t *testing.T) {
	ctx := context.Named("test-monoAPI-statement")
	err := core.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	apiKey, err := ctx.GetString(projkeys.MonoApiKey)
	if err != nil {
		t.Fatal(err)
	}

	var (
		account       = "removed"
		from    int64 = 1716498000
		to      int64 = 1716562892
	)

	url, err := ctx.GetString(projkeys.MonoApiUrl)
	if err != nil {
		t.Fatal(err)
	}

	mbClient := monobank.NewClient(url, apiKey)

	got, err := mbClient.Statement(ctx, account, from, to)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(got)
}
