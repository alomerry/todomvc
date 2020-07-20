package service

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"todomvc/core/dao"
	"todomvc/proto/todo"
)

func (*TodoService) UpdateTodo(ctx context.Context, request *todo.UpdateTodoRequest) (*todo.EmptyResponse, error) {
	fmt.Printf("request[{\n\ttoken:%s\n\tcolumn:%s\n\tvalue:%\n}]\n", request.AccessToken, request.Fields)
	// todo检查参数正确性

	err := dao.DB.UpdateId("todo", bson.ObjectIdHex(request.Id),
		bson.M{"$set": toBson(request.Fields)})
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{Message: "修改成功！"}, nil
}

func toBson(fields map[string]string) bson.M {
	res := bson.M{}

	for k, v := range fields {
		if k == "status" {
			var err error
			res[k], err = strconv.ParseBool(v)
			if err != nil {
				panic(err)
			}
			res["doneAt"] = time.Now().Unix()
		} else {
			res[k] = v
		}
	}
	return res
}
