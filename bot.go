package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"strings"
)

var bot *tgbotapi.BotAPI
var updates tgbotapi.UpdatesChannel

func init(){
	bot, _ = tgbotapi.NewBotAPI(os.Getenv("SCHEDULE_BOT"))
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ = bot.GetUpdatesChan(u)
}

func commands(update *tgbotapi.Update, msg *tgbotapi.MessageConfig){

	switch update.Message.Command() {
	case "help":
		collection := client.Database("schedule").Collection("info")

		var result Info

		filter := bson.D{}
		err := collection.FindOne(context.Background(), filter).Decode(&result)

		if err != nil { log.Panic(err) }

		msg.Text = result.Text
	case "test":
		msg.Text = string(len(strings.Split(update.Message.Text, "\n")))
	}
}

func run() {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		commands(&update, &msg)

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

type Info struct {
	Text string
}

type Subject struct {
	Day int
	Name string
	Teacher string
	Number string
	Time string
}