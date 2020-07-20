package service

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"time"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/proto/todo"
)

func (*TodoService) AddTodo(ctx context.Context, request *todo.AddTodoRequest) (*todo.EmptyResponse, error) {
	fmt.Printf("request[{\n\ttoken:%s\n\tcolor:%s\n\tcontent:%\n}]\n", request.AccessToken, request.Color, request.Content)

	userId := bson.ObjectIdHex(utils.GetJWTClaims(request.AccessToken, "jti"))
	// todo检查参数正确性
	err := dao.DB.InsertOne("todo", bson.M{
		"belongTo":  userId,
		"createdAt": time.Now().Unix(),
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
