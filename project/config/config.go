package config

import (
	"e-mobile/models"

	"github.com/ilyakaznacheev/cleanenv"
)

func LoadConfig() (models.Config, error) {
	var cfg models.Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		return models.Config{}, err
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return models.Config{}, err
	}
	return cfg, nil
}
