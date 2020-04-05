package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

// Бот.
var bot *tgbotapi.BotAPI

// Запросы.
var updates tgbotapi.UpdatesChannel

// Создание связи Telegram бота с проектом.
func init() {
	bot, _ = tgbotapi.NewBotAPI(os.Getenv("SCHEDULE_BOT"))
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ = bot.GetUpdatesChan(u)
}

// Запуск бота в Telegram.
func Run() {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		router(&update, &msg) // Маршрутизатор всех команд

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
