package main

import (
	"MyBalance/internal/context"
	"MyBalance/internal/core"
	"MyBalance/internal/telegram"
	"MyBalance/internal/telegram/functions"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	Version string
	Tag     string
	ctx     context.Context
)

func Init() error {
	context.Init(context.Config{
		Version: Version,
		Tag:     Tag,
	})

	ctx = context.Named("Init")

	if err := core.Init(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	err := Init()
	if err != nil {
		panic(err)
	}

	b, err := telegram.NewBot(ctx)
	if err != nil {
		return
	}

	b.HandleDefault("/ping", func(c tele.Context) error {
		return c.Send("Ping!")
	})

	b.Handle("/balance", functions.Balance)

	log.Println("Bot started")

	b.Start()
}
