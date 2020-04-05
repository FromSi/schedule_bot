package bot

import (
	"github.com/fromsi/schedule_bot/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

// Маршрутизатор для команд.
func router(update *tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	switch update.Message.Command() {
	case "info": // Информация
		msg.Text = info()
	case "schedule": // Расписание
		msg.Text = schedule()
	}

	// Если ничего нет, то...
	if len(strings.TrimSpace(msg.Text)) == 0 {
		msg.Text = "Нет данных"
	}
}

// Обработчик "/info" (Информация).
func info() string {
	return db.Info()
}

// Обработчик "/schedule" (Расписание).
func schedule() string {
	return db.Schedule()
}