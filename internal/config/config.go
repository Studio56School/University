package config

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func NewAppConfig() (*Config, error) {
	conf := &Config{
		"db-university.coly66w7nwap.eu-central-1.rds.amazonaws.com",
		"5432",
		"postgres",
		"12345678",
		"postgres",
	}
	return conf, nil
}
