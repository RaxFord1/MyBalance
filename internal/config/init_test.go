package config_test

import (
	"MyBalance/internal/config"
	"MyBalance/internal/context"
	"MyBalance/internal/projkeys"
	"MyBalance/internal/utils/secret"
	"testing"
)

func TestInit(t *testing.T) {
	ctx := context.Named("test-configInit")
	if err := config.Init(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log(config.GetDeploymentType(ctx))

	keys := ctx.GetKeys()
	for i, key := range keys {
		val, _ := ctx.Get(key)
		valStr, ok := val.(string)
		if !ok {
			t.Log(i, key, val)
		} else {
			t.Log(i, key, secret.ApplyMask(valStr))
		}
	}

	ctx.Get(projkeys.TelegramBotToken)
}
