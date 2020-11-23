package logic

import (
	"flag"
	"fmt"
	"gitee.com/zhuyuanhan/zyhutil/zlog"
	"github.com/pkg/errors"
	"telegram_robot/adapter"
)

type Service struct {
	configCenter ConfigCenter
	telegramBot TelegramBot
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Init(args *cmdArgs) error {
	fun := "Service.Init -->"
	var err error
	s.configCenter, err = createConfigCenter(args.configPath)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	err = initLogger(s.configCenter)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	s.telegramBot, err = createTelegramBot(s.configCenter)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	return nil
}

func (s *Service) Run() error {
	fun := "Service.Run -->"
	cmdArgs, err := s.parseFlag()
	if err != nil {
		return errors.Wrap(err, fun)
	}
	err = s.Init(cmdArgs)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	err = s.telegramBot.Run()
	if err != nil {
		return errors.Wrap(err, fun)
	}
	for {
		update := s.telegramBot.GetUpdate()
		fmt.Println(update)
		if update.Message != nil {
			s.telegramBot.SendMessage(update.Message.Chat.ID, update.Message.Text)
		}
	}
}

func createConfigCenter(path string) (ConfigCenter, error) {
	configCenterImpl := adapter.NewIniConfigCenterImpl(path)
	err := configCenterImpl.Init()
	if err != nil {
		return nil, err
	}
	return configCenterImpl, nil
}

func initLogger(configCenter ConfigCenter) error {
	normalConfig, err := configCenter.GetNormalConfig()
	if err != nil {
		return err
	}
	zlog.InitLogger(normalConfig.LogDir, "service.log", normalConfig.LogLevel)
	return nil
}

func createTelegramBot(configCenter ConfigCenter) (TelegramBot, error) {
	fun := "createTelegramBot -->"
	telegramBotConfig, err := configCenter.GetTelegramConfig()
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	telegramBot := adapter.NewTelegramBotImpl(telegramBotConfig)
	err = telegramBot.Init()
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	return telegramBot, nil
}

type cmdArgs struct {
	configPath string
}

func (s *Service) parseFlag() (*cmdArgs, error) {
	var configPath string
	flag.StringVar(&configPath, "config path", "./service.ini", "config path")
	res := &cmdArgs{
		configPath:    configPath,
	}
	return res, nil
}

