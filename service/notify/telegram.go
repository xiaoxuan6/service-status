package notify

import (
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func NewTelegram(env *Env) notify.Notifier {
	tg, _ := telegram.New(env.Telegram.Token)

	tg.AddReceivers(env.Telegram.ChatIds...)

	return tg
}
