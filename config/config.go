package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"Reward/common"
	"Reward/log"
)

type Config struct {
	FileName string
}

func Init(filename string) error {
	c := Config{
		FileName: filename,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	viper.SetConfigType("toml") // 设置配置文件格式为TOML
	if c.FileName != common.StringEmpty {
		viper.SetConfigFile(c.FileName) // TODO仅传入文件名还是还需要传入一个路径
	} else {
		viper.SetConfigFile("conf/config.toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}
