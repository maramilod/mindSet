package config

import (
	. "mind-set/config/autoload"
	utils "mind-set/internal/utils"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Conf struct {
	AppConfig `mapstructure:"app"`
	Mysql     MysqlConfig    `mapstructure:"mysql"`
	Redis     RedisConfig    `mapstructure:"redis"`
	Logger    LoggerConfig   `mapstructure:"logger"`
	Jwt       JwtConfig      `mapstructure:"jwt"`
	Firebase  FirebaseConfig `mapstructure:"firebase"`
}

var (
	Config = &Conf{
		AppConfig: App,
		Mysql:     Mysql,
		Redis:     Redis,
		Logger:    Logger,
		Jwt:       Jwt,
		Firebase:  Firebase,
	}
	once sync.Once
	V    *viper.Viper
)

func InitConfig(configPath string) {
	once.Do(func() {
		//  .yaml
		load(configPath)
		checkJwtSecretKey()
	})
}

// checkJwtSecretKey
func checkJwtSecretKey() {
	//secretKey
	if Config.Jwt.SecretKey == "" {
		Config.Jwt.SecretKey = utils.RandString(64)
		V.Set("jwt.secret_key", Config.Jwt.SecretKey)
		err := V.WriteConfig()
		if err != nil {
			panic("自动生成JWT secretKey失败: " + err.Error())
		}
	}
}

func load(configPath string) {

	var filePath string
	if configPath == "" {
		runDirectory := utils.GetRunPath()
		filePath = filepath.Join(runDirectory, "/config.yaml")
	} else {
		filePath = configPath
	}
	V = viper.New()
	V.SetConfigFile(filePath)

	if err := V.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config File Not Found: " + err.Error())
		} else {
			panic(err.Error())
		}
	}

	if err := V.Unmarshal(&Config); err != nil {
		panic(err)
	}

	if Config.WatchConfig {
		V.WatchConfig()
		V.OnConfigChange(func(in fsnotify.Event) {
			if err := V.ReadInConfig(); err != nil {
				panic(err)
			}
			if err := V.Unmarshal(&Config); err != nil {
				panic(err)
			}
		})
	}
}
