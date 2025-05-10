package handlers

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strings"
)

// HandleUpdate — маршрутизатор апдейтов
func HandleUpdate(bot *telego.Bot, upd *telego.Update) {
	if upd.Message == nil || upd.Message.Text == "" {
		return
	}

	switch parseCommand(upd.Message.Text) {
	case "start":
		onStart(bot, upd.Message)
	case "help":
		onHelp(bot, upd.Message)
	default:
		onEcho(bot, upd.Message)
	}
}

func onStart(bot *telego.Bot, msg *telego.Message) {
	sentMessage, err := bot.SendMessage(context.Background(),
		tu.Message(
			tu.ID(msg.Chat.ID),
			"Привет! Я простой бот на Go. Напиши /help, чтобы узнать команды.",
		),
	)
	if err != nil {
		fmt.Printf("Ошибка при отправке start: %v\n", err)
	}
	fmt.Printf("Sent Message: %v\n", sentMessage)
}

func onHelp(bot *telego.Bot, msg *telego.Message) {
	_, err := bot.SendMessage(context.Background(),
		tu.Message(
			tu.ID(msg.Chat.ID),
			"/start — запустить бота\n/help — помощь по командам\n<любой текст> — эхо-ответ",
		),
	)
	if err != nil {
		fmt.Printf("Ошибка при отправке help: %v\n", err)
	}
}

func onEcho(bot *telego.Bot, msg *telego.Message) {
	_, err := bot.SendMessage(context.Background(),
		tu.Message(
			tu.ID(msg.Chat.ID),
			"Вы написали: "+msg.Text,
		),
	)
	if err != nil {
		fmt.Printf("Ошибка при отправке эхо: %v\n", err)
	}
}

func parseCommand(text string) string {
	if strings.HasPrefix(text, "/") {
		fields := strings.Fields(text)
		cmd := strings.SplitN(fields[0], "@", 2)[0]
		return strings.TrimPrefix(cmd, "/")
	}
	return ""
}
