package telegram

import (
	"gopkg.in/telebot.v3"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle(telebot.OnText, t.start)
}

func (t *Telegram) start(c telebot.Context) error {
	return c.Reply("HELLO")
}
