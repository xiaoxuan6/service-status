package notify

import (
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/mail"
)

func NewEmail(env *Env) notify.Notifier {
	email := mail.New(env.Email.SenderAddress, env.Email.SmtpHostAddress)
	return email
}
