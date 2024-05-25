package logger

import (
	"MyBalance/internal/http/context"
	"log"
)

func PrintInfo(ctx context.Context, msg string) {
	log.Println(msg)
}
