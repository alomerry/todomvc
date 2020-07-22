package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id"`
	UserId    bson.ObjectId `bson:"userId"`
	CreatedAt int64         `bson:"createdAt"`
	DoneAt    int64         `bson:"doneAt"`
	Color     string        `bson:"color"`
	Status    int           `bson:"status"`
	Content   string        `bson:"content"`
}
