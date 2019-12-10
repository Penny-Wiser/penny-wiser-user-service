package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	RawId     *primitive.ObjectID
	Id        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	UpdatedAt int64
}
