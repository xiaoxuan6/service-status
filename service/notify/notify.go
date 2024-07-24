package notify

import (
	"context"
	"fmt"
	"github.com/nikoksr/notify"
)

type Notify struct {
	channel notify.Notifier
}

func NewNotify(env *Env) *Notify {

	var service notify.Notifier
	switch env.Channel {
	case "dingding":
		service = NewDingding(env)
	case "email":
		service = NewEmail(env)
	case "lark":
		service = NewLark(env)
	case "telegram":
		service = NewTelegram(env)
	case "wechat":
		service = NewWechat(env)
	default:
		service = NewDefault()
	}

	return &Notify{
		channel: service,
	}
}

func (n Notify) Send(subject, message string) {
	notify.UseServices(n.channel)

	err := notify.Send(context.Background(), subject, message)
	if err != nil {
		fmt.Println("send fail: ", err.Error())
	}
}
