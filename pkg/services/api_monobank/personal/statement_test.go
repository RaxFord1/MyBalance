package personal_test

import (
	"MyBalance/internal/context"
	"MyBalance/internal/core"
	"MyBalance/internal/projkeys"
	"MyBalance/pkg/services/api_monobank/personal"
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

	got, err := personal.Statement(ctx, apiKey, account, from, to)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(got)
}
