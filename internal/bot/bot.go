package bot

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"testobot/internal/config"
	//"testobot/internal/logger"
	"github.com/mymmrac/telego"
)

// App держит в себе бот, логгер и контекст
type App struct {
	Bot    *telego.Bot
	Logger *logrus.Logger
	Ctx    context.Context
}

// NewApp создаёт экземпляр App
func NewApp(cfg *config.Config, log *logrus.Logger) (*App, error) {
	// Инициализируем telego.Bot
	b, err := telego.NewBot(cfg.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}

	return &App{
		Bot:    b,
		Logger: log,
		Ctx:    context.Background(),
	}, nil
}

// StartPolling запускает Long Polling и обрабатывает апдейты
func (a *App) StartPolling(handler func(*telego.Update)) error {
	updates, err := a.Bot.UpdatesViaLongPolling(context.Background(), &telego.GetUpdatesParams{})
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	go func() {
		for upd := range updates {
			handler(&upd)
		}
	}()
	a.Logger.Info("🟢 Bot started polling")
	// Блокируем main, чтобы программа не завершилась
	select {}
}
