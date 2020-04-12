package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
)

// Бот.
var bot *tgbotapi.BotAPI

// Запросы.
var updates tgbotapi.UpdatesChannel

// Создание связи Telegram бота с проектом.
func init() {
	var test bool = true // true -> тестовый бот
	bot, _ = tgbotapi.NewBotAPI(os.Getenv("SCHEDULE_BOT"))
	bot.Debug = test

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	if test {
		updates, _ = bot.GetUpdatesChan(u)
	} else {
		bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://"+os.Getenv("IP_SERVER")+":"+os.Getenv("BOT_PORT")+"/"+bot.Token, "/go/src/github.com/fromsi/schedule_bot/cert.pem"))
		updates = bot.ListenForWebhook("/" + bot.Token)
		log.Fatal(http.ListenAndServeTLS(":"+os.Getenv("BOT_PORT"), "/go/src/github.com/fromsi/schedule_bot/cert.pem", "/go/src/github.com/fromsi/schedule_bot/key.pem", nil))
	}
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
