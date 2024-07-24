package notify

type Env struct {
	Channel  string `json:"channel"`
	Dingding struct {
		Token  string `json:"token"`
		Secret string `json:"secret"`
	} `json:"dingding"`
	Email struct {
		SenderAddress   string `json:"sender_address"`
		SmtpHostAddress string `json:"smtp_host_address"`
	} `json:"email"`
	Lark struct {
		WebhookUrl string `json:"webhook_url"`
	} `json:"lark"`
	Telegram struct {
		Token   string  `json:"token"`
		ChatIds []int64 `json:"chat_ids"`
	} `json:"telegram"`
	Wechat struct {
		AppId          string `json:"app_id"`
		AppSecret      string `json:"app_secret"`
		Token          string `json:"token"`
		EncodingAesKey string `json:"encoding_aes_key"`
		UserIds        []string `json:"user_ids"`
	} `json:"wechat"`
}
