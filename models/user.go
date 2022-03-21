package models

import (
	"path"
	"time"
)

type UserData struct {
	Id        string      `bson:"user_id"`
	Password  string      `bson:"password"`
	IsActive  bool        `json:"is_active" bson:"is_active"`
	Balance   string      `bson:"balance"`
	Age       interface{} `json:"age,string" bson:"age"`
	Name      string      `bson:"name"`
	Gender    string      `bson:"gender"`
	Company   string      `bson:"company"`
	Email     string      `bson:"email"`
	Phone     string      `bson:"phone"`
	Address   string      `bson:"address"`
	About     string      `bson:"about"`
	Registred time.Time   `bson:"registred"`
	Latitude  float32     `bson:"latitude"`
	Longitude float32     `bson:"longitude"`
	Tags      []string    `bson:"tags"`
	Friends   []Friend    `bson:"friends"`
	Data      string      `bson:"data"`
}

type Friend struct {
	Id   int
	Name string
}

var UserDataFilePath string

func DataPathString(UserId string) string {
	return path.Join(UserDataFilePath, UserId)
}
