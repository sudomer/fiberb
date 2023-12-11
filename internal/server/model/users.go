package model

import (
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExist error = errors.New("user already exist")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username" json:"username" validate:"required"`
	Name     string             `bson:"name" json:"name" validate:"required"`
	Surname  string             `bson:"surname" json:"surname" validate:"required"`
	Email    string             `bson:"email" json:"email" validate:"required"`
	Password string             `bson:"password" json:"password" validate:"required"`
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
	res := DB_USERS.FindOne(context.Background(), bson.M{"username": username})
	res.Decode(&usr)

	return usr, nil
}

func DeleteUser(username string) (int, error) {
	res, err := DB_USERS.DeleteOne(context.Background(), bson.D{{"username", username}})
	if err != nil {
		return 0, err
	}

	return int(res.DeletedCount), nil
}

func (usr User) CreateUser() (string, error) {

	count, err := DB_USERS.CountDocuments(context.Background(), bson.M{"username": usr.Username, "email": usr.Email})
	if err != nil {
		return "", err
	}
	if count > 0 {
		return strconv.Itoa(int(count)), ErrUserAlreadyExist
	}
	temp := usr
	resp, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 8)
	if err != nil {
		return "Password encrypting process failed.", err
	}
	temp.Password = string(resp)
	res, err := DB_USERS.InsertOne(context.Background(), temp)
	if err != nil {
		return "Database user creation process failed.", err
	}
	response := res.InsertedID.(primitive.ObjectID)
	return response.Hex(), nil

}
