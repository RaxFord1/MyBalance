package main

import (
	"MyBalance/internal/context"
	"MyBalance/pkg/core"
	"MyBalance/pkg/mono_balance"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
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

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/balance", func(c tele.Context) error {
		balance, err := mono_balance.GetBalance(ctx)
		if err != nil {
			return c.Send(err.Error())
		}
		return c.Send(balance)
	})

	log.Println("Bot started")

	b.Start()
}
