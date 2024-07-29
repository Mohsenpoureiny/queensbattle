package telegram

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"queensbattle/internal/service"
	"queensbattle/internal/telegram/teleprompt"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, apiKey string) (*Telegram, error) {
	t := &Telegram{
		App:        app,
		TelePrompt: teleprompt.NewTelePrompt(),
	}

	pref := telebot.Settings{
		Token:   apiKey,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: t.onError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Fatal("couldn't connect to telegram bot")
		return nil, err
	}
	t.bot = bot

	t.setupHandlers()
	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}
