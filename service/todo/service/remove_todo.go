package service

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/proto/todo"
)

func (*TodoService) RemoveTodo(ctx context.Context, request *todo.RemoveTodoRequest) (*todo.EmptyResponse, error) {
	fmt.Printf("request[{\n\ttoken:%s\n\tid:%s\n}]\n", request.AccessToken, request.Id)

	// todo检查参数正确性
	err := dao.DB.RemoveId("todo", bson.ObjectIdHex(request.Id))
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{Message: "删除成功！"}, nil
}
