package notify

import (
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/lark"
)

func NewLark(env *Env) notify.Notifier {
	lark := lark.NewWebhookService(env.Lark.WebhookUrl)

	return lark
}
