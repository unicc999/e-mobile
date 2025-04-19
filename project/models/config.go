package models

type Config struct {
	DatabaseHost     string `env:"DB_HOST"`
	DatabaseName     string `env:"DB_NAME"`
	DatabasePassword string `env:"DB_PASSWORD"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePort     string `env:"DB_PORT"`
}
