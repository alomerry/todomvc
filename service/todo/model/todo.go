package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id         bson.ObjectId `bson:"_id"`
	CreatedAt  time.Time     `bson:"createdAt"`
	FinishedAt time.Time     `bson:"finishedAt"`
	Status     string        `bson:"status"`
	Content    string        `bson:"content"`
}
