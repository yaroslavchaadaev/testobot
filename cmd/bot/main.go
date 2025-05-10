package main

import (
	"github.com/mymmrac/telego"
	"log"
	"testobot/internal/bot"
	"testobot/internal/config"
	"testobot/internal/handlers"
	"testobot/internal/logger"
)

func main() {
	// 1) Загружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ cannot load config: %v", err)
	}

	// 2) Настраиваем логгер
	logg := logger.New()

	// 3) Создаём приложение (бот + логгер)
	app, err := bot.NewApp(cfg, logg)
	if err != nil {
		logg.Fatalf("❌ cannot create bot: %v", err)
	}

	// 4) Запускаем polling с нашим хендлером
	if err := app.StartPolling(func(upd *telego.Update) {
		handlers.HandleUpdate(app.Bot, upd)
	}); err != nil {
		logg.Fatalf("❌ polling failed: %v", err)
	}
}
