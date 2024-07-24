package notify

type Env struct {
	Channel  string `yaml:"channel"`
	Dingding struct {
		Token  string `yaml:"token"`
		Secret string `yaml:"secret"`
	} `yaml:"dingding"`
	Email struct {
		SenderAddress   string `yaml:"sender_address"`
		SmtpHostAddress string `yaml:"smtp_host_address"`
	} `yaml:"email"`
	Lark struct {
		WebhookUrl string `yaml:"webhook_url"`
	} `yaml:"lark"`
	Telegram struct {
		Token   string  `yaml:"token"`
		ChatIds []int64 `yaml:"chat_ids"`
	} `yaml:"telegram"`
	Wechat struct {
		AppId          string `yaml:"app_id"`
		AppSecret      string `yaml:"app_secret"`
		Token          string `yaml:"token"`
		EncodingAesKey string `yaml:"encoding_aes_key"`
		UserIds        []string `yaml:"user_ids"`
	} `yaml:"wechat"`
}
