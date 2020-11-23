package adapter

import (
	"gitee.com/zhuyuanhan/zyhutil/zconfig"
	"github.com/pkg/errors"
	"telegram_robot/util"
)

const (
	SectionNormal = "normal"
	SectionMysql = "mysql"
	SectionTelegram = "telegram_bot"
)

type IniConfigCenterImpl struct {
	configPath string
	cfg *zconfig.ConfigReader
}

func NewIniConfigCenterImpl(configPath string) *IniConfigCenterImpl {
	return &IniConfigCenterImpl{
		configPath: configPath,
	}
}

func (i *IniConfigCenterImpl) Init() error {
	fun := "IniConfigCenterImpl.Init -->"
	cfg, err := zconfig.LoadConfig(i.configPath)
	if err != nil {
		return errors.Wrap(err, fun)
	}
	i.cfg = cfg
	return nil
}

func (i *IniConfigCenterImpl) GetNormalConfig() (*util.NormalConfig, error) {
	fun := "IniConfigCenterImpl.GetNormalConfig -->"
	res := &util.NormalConfig{}
	err := i.cfg.UnmarshalWithSection(SectionNormal, res)
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	return res, nil
}

func (i *IniConfigCenterImpl) GetMysqlConfig() (*util.MysqlConfig, error) {
	fun := "IniConfigCenterImpl.GetMysqlConfig -->"
	res := &util.MysqlConfig{}
	err := i.cfg.UnmarshalWithSection(SectionMysql, res)
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	return res, nil
}

func (i *IniConfigCenterImpl) GetTelegramConfig() (*util.TelegramBotConfig, error) {
	fun := "IniConfigCenterImpl.TelegramBotConfig -->"
	res := &util.TelegramBotConfig{}
	err := i.cfg.UnmarshalWithSection(SectionTelegram, res)
	if err != nil {
		return nil, errors.Wrap(err, fun)
	}
	return res, nil
}