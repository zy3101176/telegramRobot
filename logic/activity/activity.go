package activity

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Activity interface {
	Init() error
	DataProcess(status int32, update *tgbotapi.Update) (int32, *tgbotapi.MessageConfig, error)
	IsFinish(status int32) bool
}
