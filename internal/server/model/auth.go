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
	err := DB_USERS.FindOne(context.Background(), bson.M{"username": auth.Username}).Decode(&usr)
	if err != nil {
		return ErrUserNotFound, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(auth.Password))
	if err != nil {
		return ErrWrongPassword, err
	}
	return usr.Email, nil
}
