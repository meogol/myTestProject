package config

import (
	"encoding/json"
	"errors"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Database DataBaseConfig
}

type DataBaseConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	SslMode    string
}

var configPath = "./configs"
var configName = "config"

// ---
// here are defaults
var CurrentConfig = &AppConfig{
	Database: DataBaseConfig{DbHost: "localhost", DbUser: "user", DbPassword: "password", DbName: "db"},
}

func ReloadConfig() {
	defer func() {
		jsonConfig, _ := json.Marshal(CurrentConfig)
		configLogger.Debugf("result config is", "config", string(jsonConfig))
	}()

	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			configLogger.DPanicf("failed to read config %s", err)
		}
		configLogger.Debugf("not found file with name %s in path %s; using default config", configName, configPath)
		return
	}
	err = viper.Unmarshal(&CurrentConfig)
	if err != nil {
		configLogger.DPanicf("failed to parse config %s", err)
	}
}

// ---
func init() {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	ReloadConfig()
}
