package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config хранит параметры из окружения
type Config struct {
	BotToken string
}

// Load читает .env (если есть) и возвращает Config
func Load() (*Config, error) {
	// Попытаемся загрузить .env, но без фатала, если его нет
	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("BOT_TOKEN is not set in env")
	}

	return &Config{BotToken: token}, nil
}
