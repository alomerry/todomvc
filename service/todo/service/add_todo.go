package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"time"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/proto/todo"
)

func (*TodoService) AddTodo(ctx context.Context, request *todo.AddTodoRequest) (*todo.EmptyResponse, error) {

	userId := bson.ObjectIdHex(utils.GetJWTClaims(request.AccessToken, "jti"))
	err := dao.DB.InsertOne("todo", bson.M{
		"belongTo":  userId,
		"createdAt": time.Now().UnixNano() / 1000,
		"doneAt":    0,
		"color":     request.Color,
		"content":   request.Content,
		"status":    false,
	})
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{Message: "新建成功！"}, nil
}
