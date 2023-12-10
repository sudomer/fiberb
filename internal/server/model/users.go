package model

import (
	"context"
	"errors"

	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty" validate:"required"`
	Name     string             `bson:"name" validate:"required"`
	Surname  string             `bson:"surname" validate:"required"`
	Email    string             `bson:"email" validate:"required"`
	Password string             `bson:"password" validate:"required"`
	Contact  Contact            `bson:"contact,omitempty,inline"`
}

type Contact struct {
	Phone   string   `bson:"phone,omitempty"`
	Emails  []string `bson:"extra_emails,omitempty"`
	Discord string   `bson:"discord,omitempty"`
}

func GetUsers() ([]User, error) {
	cr, err := DB_USERS.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cr.Close(context.TODO())

	var usr_data []User
	for cr.Next(context.Background()) {
		var usr User
		if err := cr.Decode(&usr); err != nil {
			lib.Log().Error("User data can't migrate", zap.Error(err))
			return nil, err
		}
		tempUsers := User{
			ID:       usr.ID,
			Username: usr.Username,
			Name:     usr.Name,
			Surname:  usr.Surname,
			Email:    usr.Email,
		}

		usr_data = append(usr_data, tempUsers)
	}
	return usr_data, nil
}

func GetUser(username string) (User, error) {
	var usr User
	cr, err := DB_USERS.Find(context.Background(), bson.D{{Key: "username", Value: username}})
	if err != nil {
		lib.Log().Error("User data can't found", zap.Error(err))
		return User{}, err
	}
	cr.Decode(&usr)

	return usr, nil
}

func DeleteUser(_id string) error {
	res, err := DB_USERS.DeleteOne(context.Background(), bson.M{"_id": _id})
	if err != nil {
		lib.Log().Error("User data can not be deleted", zap.Error(err))
		return err
	}
	lib.Log().Error("User data deleted", zap.Any("deleted count", res.DeletedCount))
	return nil
}

func (usr User) CreateUser() (string, error) {

	// Username check

	count, err := DB_USERS.CountDocuments(context.Background(), bson.M{"username": usr.Username, "email": usr.Email})
	if err != nil {
		lib.Log().Error("User data can't found", zap.Error(err))
		return "", err
	}
	lib.Log().Warn("Count", zap.Int64("count", count))
	if count > 0 {
		return "User already exist", errors.New("user already exist")
	}

	temp := usr
	resp, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 8)
	if err != nil {
		lib.Log().Error("User can not be created", zap.Error(err))
		return "", err
	}
	temp.Password = string(resp)
	res, err := DB_USERS.InsertOne(context.Background(), temp)
	if err != nil {
		lib.Log().Error("User can not be created", zap.Error(err))
		return "", err
	}
	response := res.InsertedID.(primitive.ObjectID)
	return response.Hex(), nil

}
