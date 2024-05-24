package telegram

import (
	"MyBalance/internal/projkeys"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

func (b *Bot) Handle(endpoint interface{}, h HandleFunc, m ...tele.MiddlewareFunc) {
	wrappedHandler := func(c tele.Context) error {
		b.ctx.Set(projkeys.ClientID, strconv.FormatInt(c.Update().Message.Sender.ID, 10))
		return h(b.ctx, c)
	}
	b.bot.Handle(endpoint, wrappedHandler, m...)
}

func (b *Bot) HandleDefault(endpoint interface{}, h tele.HandlerFunc, m ...tele.MiddlewareFunc) {
	b.bot.Handle(endpoint, h, m...)
}
