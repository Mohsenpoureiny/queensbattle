package telegram

import (
	"gopkg.in/telebot.v3"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, t.textHandler)
}

func (t *Telegram) textHandler(context telebot.Context) error {
	if t.TelePrompt.Dispatch(context.Sender().ID, context) {
		return nil
	}

	return context.Reply("I didn't understand")
}
