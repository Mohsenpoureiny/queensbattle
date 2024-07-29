package telegram

import (
	"errors"
	"gopkg.in/telebot.v3"
)

var (
	ErrInputTimeout = errors.New("input timeout")
)

type InputConfig struct {
	Prompt    any
	OnTimeout any
}

func (t *Telegram) Input(c telebot.Context, config InputConfig) (*telebot.Message, error) {

	if config.Prompt != nil {
		c.Reply(config.Prompt)
	}

	response, isTimeout := t.TelePrompt.AsMessage(c.Sender().ID, DefaultInputTimeout)
	if isTimeout {
		if config.OnTimeout != nil {
			c.Reply(config.OnTimeout)
		}
		return nil, ErrInputTimeout
	}

	return response, nil
}
