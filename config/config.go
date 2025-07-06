package config

import (
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config/app.yaml"
	}
	return getConfig(path)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
