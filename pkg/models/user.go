package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" 			json:"id"`
	Name        string        `bson:"name" 			json:"name"`
	UserName    string        `bson:"username" 		json:"username"`
	Email       string        `bson:"email"			json:"email"`
	BirthDate   time.Time     `bson:"birthDate"		json:"birthDate"`
	Phone       string        `bson:"phone" 		json:"phone"`
	Description string        `bson:"description" 	json:"description"`
}

type Users []User

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
