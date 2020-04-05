package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Текущий клиент.
var client *mongo.Client

// Название БД.
const dbName string = "schedule"

// Путь до mongodb.
const mongoUri string = "mongodb://localhost:27017"

// Создание клиента, чтобы работать с БД.
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
}
