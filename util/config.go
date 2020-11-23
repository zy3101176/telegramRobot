package util

type NormalConfig struct {
	LogLevel  string `ini:"log_level"`
	LogDir string `ini:"log_dir"`
}

type TelegramBotConfig struct {
	HttpProxy string  `ini:"http_proxy"`
	Token     string  `ini:"telegram_token"`
	UpdateTimeout int `ini:"update_timeout"`
	IsDebug   bool    `ini:"is_debug"`
}

type MysqlConfig struct {
	Address  string `ini:"address"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
