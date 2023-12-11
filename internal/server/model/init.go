package model

import (
	"context"
	"os"

	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Client = *mongo.Client

var (
	DB       Client
	DB_USERS *mongo.Collection
)

func init() {
	// Connection on database
	mongoURI := os.Getenv("DATABASE_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		lib.Log().Panic("Database connection error", zap.Error(err))
	}
	res, err := client.Database("sudomer").Collection("users").CountDocuments(context.Background(), bson.D{})
	if err != nil {
		lib.Log().Panic("Database collection error", zap.Error(err))
	}
	lib.Log().Info("Database connected.", zap.Int("Members", int(res)))
	DB = client
	DB_USERS = client.Database("sudomer").Collection("users")
}
