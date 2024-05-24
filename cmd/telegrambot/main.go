package main

import (
	"MyBalance/internal/config"
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

	// different config types for different deployment types
	config.SetConfigType(ctx, config.TypeOfConfigFromFile)
	config.SetDeploymentInfoSource(ctx, config.DeploymentInfoFromEnv)

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
		panic(err)
	}

	b.HandleDefault("/start", func(c tele.Context) error {
		return c.Send("Привет! Чтобы узнать баланс, напиши /balance. " +
			"Потом можешь написать /statement, чтобы узнать историю транзакций за текущий день. \n" +
			"В конце дня тебе будет приходить текущий баланс и история транзакций.")
	})

	b.Handle("/balance", functions.Balance)

	b.Handle("/statement", functions.Statement)

	b.HandleDefault("/ping", func(c tele.Context) error {
		return c.Send("Ping!")
	})

	log.Println("Bot started")

	b.Start()
}
