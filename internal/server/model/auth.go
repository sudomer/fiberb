package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Username string `bson:"username" json:"username" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
}

const (
	ErrUserNotFound  string = "user not found"
	ErrWrongPassword string = "wrong password"
)

func (auth Auth) Login() (string, error) {
	var usr User
	err := DB_USERS.FindOne(context.TODO(), bson.M{"username": auth.Username}).Decode(&usr)
	if err != nil {
		return ErrUserNotFound, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(auth.Password))
	if err != nil {
		return ErrWrongPassword, err
	}
	return "success", nil
}

func (auth Auth) Register(usr User) (string, error) {
	usr.Username = auth.Username
	usr.Password = auth.Password

	resp, err := usr.CreateUser()
	if err != nil {
		return "user not created", err
	}

	return resp, nil
}
