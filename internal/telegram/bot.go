package telegram

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	tele "gopkg.in/telebot.v3"
	"time"
)

type IBot interface {
	Handle(endpoint interface{}, h HandleFunc, m ...tele.MiddlewareFunc)
	HandleDefault(endpoint interface{}, h tele.HandlerFunc, m ...tele.MiddlewareFunc)
	Start()
	Send(ctx context.Context, to string, message interface{}, opts ...interface{}) (*tele.Message, error)
}
type HandleFunc func(ctx context.Context, ctx2 tele.Context) error

var _ IBot = (*Bot)(nil)

type Bot struct {
	ctx context.Context
	bot *tele.Bot
}

func (b *Bot) Send(ctx context.Context, to string, message interface{}, opts ...interface{}) (*tele.Message, error) {
	r := Recipient{chatId: to}
	return b.bot.Send(r, message, opts...)
}

func (b *Bot) Start() {
	b.bot.Start()
}

func NewBot(ctx context.Context) (*Bot, error) {
	token, err := ctx.GetString(projkeys.TelegramBotToken)
	if err != nil {
		return nil, requesto.InternalError.NewWithMsg(ctx, err.Error())
	}

	pollingTimeout := ctx.GetIntOptional(projkeys.TelegramLongPollingTimeout, 10)

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: time.Duration(pollingTimeout) * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &Bot{
		ctx: ctx,
		bot: b,
	}, nil
}

func BotFromTeleBot(ctx context.Context, bot *tele.Bot) *Bot {
	return &Bot{
		ctx: ctx,
		bot: bot,
	}
}
