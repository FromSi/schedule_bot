package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

// Текущий клиент.
var client *mongo.Client

// Название БД.
const dbName string = "schedule"

// Создание клиента, чтобы работать с БД.
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("SCHEDULE_MONGO_URL")))
}
