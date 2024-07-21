package main

import (
	"MyBalance/internal/config"
	"MyBalance/internal/core"
	"MyBalance/internal/http/context"
	"MyBalance/internal/tasks"
	"MyBalance/internal/telegram"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	config.SetConfigType(ctx, config.TypeOfConfigFromMemory)
	config.SetDeploymentInfoSource(ctx, config.DeploymentInfoFromEnv)

	if err := core.Init(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Init(); err != nil {
		log.Fatal(err)
		return
	}

	//summary, err := sendSummary(events.APIGatewayProxyRequest{})
	//log.Print(summary)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	lambda.Start(sendSummary)
}

func sendSummary(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.NewFromPrev(ctx, "aws-lambda-send-summary")

	bot, err := telegram.NewBot(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	if err = tasks.GetAndSendCurrentBalanceDaySummary(ctx, bot); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
