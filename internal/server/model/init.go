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
	res, err := client.Database("sudomer").Collection("users").Find(context.Background(), bson.D{})
	if err != nil {
		lib.Log().Panic("Database collection error", zap.Error(err))
	}
	defer res.Close(context.TODO())
	var usr_data int
	for res.Next(context.Background()) {
		var usr User
		if err := res.Decode(&usr); err != nil {
			lib.Log().Error("User data can't migrate", zap.Error(err))
		}
		usr_data += 1
	}
	lib.Log().Info("Database connected.", zap.Int("Members", usr_data))
	DB = client
	DB_USERS = client.Database("sudomer").Collection("users")
}
