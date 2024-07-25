package notify

import (
	"context"
	"fmt"
	"sync"

	"github.com/nikoksr/notify"
)

type Notify struct {
	channel notify.Notifier
	target  bool
}

var (
	notifyInstance *Notify
	once           sync.Once
	mu             sync.Mutex
)

func NewNotify(env *Env) *Notify {
	once.Do(func() {
		notifyInstance = &Notify{}
	})
	mu.Lock()
	defer mu.Unlock()
	notifyInstance.updateNotifier(env)
	return notifyInstance
}

func (n *Notify) updateNotifier(env *Env) {
	var service notify.Notifier
	target := true
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
		target = false
		service = nil
		if len(env.Channel) > 1 {
			fmt.Printf("invalid notify channel [%s]", env.Channel)
		}
	}
	n.channel = service
	n.target = target
}

func (n Notify) Send(subject, message string) {
	if n.target {
		notifyService := notify.NewWithServices(n.channel)
		err := notifyService.Send(context.Background(), subject, message)
		if err != nil {
			fmt.Println("send fail: ", err.Error())
		}
	}
}
