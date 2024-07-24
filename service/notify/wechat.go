package notify

import (
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/wechat"
	"github.com/silenceper/wechat/v2/cache"
)

func NewWechat(env *Env) notify.Notifier {
	memory := cache.NewMemory()
	w := wechat.New(&wechat.Config{
		AppID:          env.Wechat.AppId,
		AppSecret:      env.Wechat.AppSecret,
		Token:          env.Wechat.Token,
		EncodingAESKey: env.Wechat.EncodingAesKey,
		Cache:          memory,
	})

	w.AddReceivers(env.Wechat.UserIds...)

	return w
}
