package logger

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"log"
)

func PrintError(ctx context.Context, name string, err error) {
	log.Println(requesto.InternalError.NewWithMsg(ctx, name+": "+err.Error()))
}
