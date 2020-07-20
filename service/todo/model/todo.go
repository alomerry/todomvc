package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id"`
	BelongTo  bson.ObjectId `bson:"belongTo"`
	CreatedAt int64        `bson:"createdAt"`
	DoneAt    int64        `bson:"doneAt"`
	Color     string        `bson:"color"`
	Status    bool        `bson:"status"`
	Content   string        `bson:"content"`
}
