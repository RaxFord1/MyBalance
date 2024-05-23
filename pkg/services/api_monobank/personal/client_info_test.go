package personal_test

import (
	"MyBalance/internal/context"
	"MyBalance/pkg/services/api_monobank/personal"
	"log"
	"testing"
)

func TestClientInfo(t *testing.T) {
	ctx := context.Named("test-monoAPI-client-info")
	secretToken := ""

	got, err := personal.ClientInfo(ctx, secretToken)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(got)
}
