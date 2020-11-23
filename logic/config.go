package logic

import (
	"telegram_robot/util"
)

type ConfigCenter interface {
	Init() error
	GetNormalConfig() (*util.NormalConfig, error)
	GetMysqlConfig() (*util.MysqlConfig, error)
	GetTelegramConfig() (*util.TelegramBotConfig, error)
}