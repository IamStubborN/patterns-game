package config

import (
	"github.com/joho/godotenv"
	"github.com/tomazk/envcfg"
)

// Config - default app config struct
type Config struct {
	TelegramToken string `envcfg:"TELEGRAM_TOKEN"`
}

//InitConfig - get env params into struct
func InitConfig() (Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return cfg, err
	}

	if err := envcfg.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
