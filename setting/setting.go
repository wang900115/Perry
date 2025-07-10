package setting

import (
	"os"

	"github.com/spf13/viper"
)

func NewSetting() *viper.Viper {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "setting/setting.json"
	}
	return getSetting(path)
}

func getSetting(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
