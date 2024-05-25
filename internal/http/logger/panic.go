package logger

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"fmt"
	"log"
	"time"
)

func PanicAndCall(ctx context.Context, f interface{}, timeout time.Duration) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("PanicAndCall: %v", r)))
		}
	}()

	time.Sleep(timeout)

	f.(func(ctx context.Context))(ctx)
}
