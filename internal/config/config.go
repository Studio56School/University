package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func NewAppConfig() (*Config, error) {

	viper.SetConfigType("json") // Look for specific type
	filepath := "./heml/config.json"
	if os.Getenv("path") != "" {
		filepath = os.Getenv("path")
	}
	file, err := os.Open(filepath)
	if err != nil {
		panic("failed to open file with config")
	}
	err = viper.ReadConfig(file)
	if err != nil {
		panic("failed to read config")
	}

	conf := &Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.name_db"),
	}

	return conf, nil
}
