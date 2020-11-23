package adapter

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"telegram_robot/util"
)

type TelegramBotImpl struct {
	cfg *util.TelegramBotConfig
	bot  *tgbotapi.BotAPI
	updateCh tgbotapi.UpdatesChannel
}

func NewTelegramBotImpl(cfg *util.TelegramBotConfig) *TelegramBotImpl {
	return &TelegramBotImpl{
		cfg: cfg,
	}
}

func (t *TelegramBotImpl) Init() error {
	fun := "TelegramBotImpl.Init -->"
	client := getClientWithProxy(t.cfg.HttpProxy)
	bot, err := tgbotapi.NewBotAPIWithClient("1458734494:AAF3lg0guCQJKuw5h225Ohd7JTf0p80j8gM", client)
	bot.Debug = t.cfg.IsDebug
	if err != nil {
		return errors.Wrap(err, fun)
	}
	t.bot = bot
	return nil
}

func (t *TelegramBotImpl) Run() error {
	fun := "TelegramBotImpl.Run -->"
	u := tgbotapi.NewUpdate(0)
	u.Timeout = t.cfg.UpdateTimeout
	updates, err := t.bot.GetUpdatesChan(u)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	t.updateCh = updates
	return nil
}

func (t *TelegramBotImpl) GetUpdate() *tgbotapi.Update {
	fmt.Println(len(t.updateCh))
	res := <-t.updateCh
	return &res
}

func (t *TelegramBotImpl) SendMessage(chatId int64, text string) error {
	fun := "TelegramBotImpl.SendMessage -->"
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := t.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	return nil
}

func (t *TelegramBotImpl) Stop() {
	t.bot.StopReceivingUpdates()
}

func getClientWithProxy(proxyAddr string) *http.Client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1087")
	}
	transport := &http.Transport{Proxy: proxy}
	return &http.Client{Transport: transport}
}