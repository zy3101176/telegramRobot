package logic

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramBot interface {
	Init() error
	Run() error
	Stop()
	GetUpdate() *tgbotapi.Update
	SendMessage(chatId int64, text string) error
}