package mono_balance

import (
	"MyBalance/internal/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
)

func LimitCheck(ctx context.Context, key string) error {
	if err := limiterAll.Allow(key); err != nil {
		return err
	}

	clientID, found := ctx.Get(projkeys.ClientID)
	if found {
		clientIdStr, ok := clientID.(string)
		if !ok {
			return requesto.InternalError.NewWithMsg(ctx, "clientId not string")
		}
		keyClientLimit := formClientId(key, clientIdStr)
		if err := limiterClient.Allow(keyClientLimit); err != nil {
			return err
		}
	}

	return nil
}

func formClientId(key, clientID string) string {
	return key + "_" + clientID
}
