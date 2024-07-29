package telegram

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)

	account := GetAccount(c)

	if !isJustCreated {
		t.myInfo(c)
	}

	msg, err := t.Input(c, InputConfig{
		Prompt: `👋 سلام به QueenBattle خوش آمدی.
میخوای کاربرهای دیگه  به چه اسمی ببیننت؟ این اسم رو بعدا هم میتونی تغییر بدی`,
		OnTimeout: DefaultInputText,
	})

	if err != nil {
		return err
	}

	displayName := msg.Text

	// todo: displayName validation
	account.DisplayName = displayName

	if err := t.App.Account.Update(context.Background(), account); err != nil {
		return err
	}

	 c.Reply(fmt.Sprintf(`✅ از این به بعد شما را %s صدا میزنیم.`, displayName))
	 return t.myInfo(c)
}

func (t *Telegram) myInfo(c telebot.Context) error {
	account := GetAccount(c)
	return c.Reply(fmt.Sprintf(`سلام %s خوش آمدی.`, account.DisplayName)))
}
