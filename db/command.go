package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// Получить информацию.
func Info() string {
	collection := client.Database(dbName).Collection("info")

	var result info

	filter := bson.D{}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Panic(err)
	}

	return result.toString()
}

// Получить расписание на текущий день.
func Schedule() (text string) {
	collection := client.Database(dbName).Collection("schedule")

	var result schedule

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{{"weekday", int(time.Now().Weekday())}}
	cur, _ := collection.Find(ctx, filter)

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}

		text += result.toString()
	}

	return text
}