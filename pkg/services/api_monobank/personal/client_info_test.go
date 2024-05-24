package personal_test

import (
	"MyBalance/internal/context"
	"MyBalance/internal/core"
	"MyBalance/internal/projkeys"
	"MyBalance/pkg/services/api_monobank/personal"
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

	got, err := personal.ClientInfo(ctx, apiKey)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(got)
}
