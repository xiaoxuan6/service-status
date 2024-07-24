package notify

import (
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/dingding"
)

func NewDingding(env *Env) notify.Notifier {
	service := dingding.New(&dingding.Config{Token: env.Dingding.Token, Secret: env.Dingding.Secret})

	return service
}
