package config

import (
	"encoding/json"
	"errors"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Database *DataBaseConfig
	Server   *ServerConfig
}

type DataBaseConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	SslMode    string
}

type ServerConfig struct {
	Port int
	Host string
}

var configPath = "./configs"
var configName = "config"

// ---
// here are defaults
var CurrentConfig = &AppConfig{
	Database: &DataBaseConfig{DbHost: "localhost", DbUser: "user", DbPassword: "password", DbName: "db", DbPort: "5432", SslMode: "disable"},
	Server:   &ServerConfig{Port: 8080, Host: "0.0.0.0"},
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
