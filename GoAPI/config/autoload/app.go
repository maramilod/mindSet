package autoload

import (
	"mind-set/internal/utils"
	"mind-set/internal/utils/convert"
)

type AppConfig struct {
	AppEnv         string `mapstructure:"app_env"`
	Debug          bool   `mapstructure:"debug"`
	Language       string `mapstructure:"language"`
	WatchConfig    bool   `mapstructure:"watch_config"`
	StaticBasePath string `mapstructure:"base_path"`
}

var App = AppConfig{
	AppEnv:         "local", //production
	Debug:          true,    //false
	Language:       "en_US",
	WatchConfig:    false,
	StaticBasePath: getDefaultPath(),
}

func getDefaultPath() (path string) {
	path, _ = utils.GetDefaultPath()
	path = convert.GetString(utils.If(path != "", path, "/tmp"))
	return
}
