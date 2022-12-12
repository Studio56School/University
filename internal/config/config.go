package config

import "github.com/spf13/viper"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func NewAppConfig() (*Config, error) {

	conf := &Config{
		viper.GetString("db_host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.name_db"),
	}

	return conf, nil
}
