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
	Status    bool          `bson:"status"`
	Content   string        `bson:"content"`
}

func NewProtoTodo(id, userId bson.ObjectId, createdAt, doneAt int64, color string, status bool, content string) *Todo {
	return &Todo{
		Id:        id,
		UserId:    userId,
		CreatedAt: createdAt,
		DoneAt:    doneAt,
		Color:     color,
		Status:    status,
		Content:   content,
	}
}
