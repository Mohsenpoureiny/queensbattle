package telegram

import (
	"gopkg.in/telebot.v3"
	"queensbattle/internal/entity"
	"time"
)

var (
	DefaultInputTimeout = time.Minute * 5
	DefaultInputText    = `🕐 منتظر پیامت بودیم ولی چیزی ارسال نکردی، لطاق هر وقت برگشتی دوباره امتحان کن.`
)

func GetAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
