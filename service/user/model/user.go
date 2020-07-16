package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"createdAt"`
}
