package telegram

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"queensbattle/internal/entity"
)

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			Username:  c.Sender().Username,
		}
		account, created, err := t.App.Account.CreateOrUpdate(context.Background(), acc)

		if err != nil {
			return err
		}

		c.Set("account", account)
		c.Set("is_just_created", created)

		return next(c)
	}
}

func (t *Telegram) onError(err error, context telebot.Context) {
	if errors.Is(err, ErrInputTimeout) {
		return
	}

	errorId := uuid.New().String()

	logrus.WithError(err).WithField("tracing_id", errorId).Errorln("unhandled error")

	context.Reply(fmt.Sprintf("در پردازش دستور مشکلی پیش آمد \n کد برسی: %s", errorId))
}
