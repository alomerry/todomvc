package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"time"
	"todomvc/core/dao"
	"todomvc/proto/todo"
)

func (*TodoService) AddTodo(ctx context.Context, request *todo.AddTodoRequest) (*todo.EmptyResponse, error) {

	userId := bson.ObjectIdHex(request.UserId)
	err := dao.DB.InsertOne("todo", bson.M{
		"userId":    userId,
		"createdAt": time.Now().UnixNano() / 1000,
		"doneAt":    0,
		"color":     request.Color,
		"content":   request.Content,
		"status":    false,
	})
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{}, nil
}
