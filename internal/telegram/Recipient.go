package telegram

import tele "gopkg.in/telebot.v3"

var _ tele.Recipient = (*Recipient)(nil)

type Recipient struct {
	chatId string
}

func (r Recipient) Recipient() string {
	return r.chatId
}
