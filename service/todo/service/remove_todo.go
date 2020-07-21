package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/proto/todo"
)

func (*TodoService) RemoveTodo(ctx context.Context, request *todo.RemoveTodoRequest) (*todo.EmptyResponse, error) {
	err := dao.DB.RemoveById("todo", bson.ObjectIdHex(request.Id))
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{}, nil
}
