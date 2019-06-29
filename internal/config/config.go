package config

import (
	"github.com/spf13/viper"
	"sync"
)

const (
	configFileName = "config"
	configFilePath = "./config"
)

var (
	config *viper.Viper
	once   sync.Once
)

func Config() (*viper.Viper, error) {
	var err error
	once.Do(func() {
		config = viper.New()
		config.SetConfigName(configFileName)
		config.AddConfigPath(configFilePath)
		e := config.ReadInConfig()
		if e != nil {
			err = e
		}
	})
	if err != nil {
		return nil, err
	}
	return config, nil
}
