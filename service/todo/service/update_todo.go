package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"todomvc/core/dao"
	"todomvc/proto/todo"
)

func (*TodoService) UpdateTodo(ctx context.Context, request *todo.UpdateTodoRequest) (*todo.EmptyResponse, error) {
	if dao.DB.UpdateById("todo", bson.ObjectIdHex(request.Id), bson.M{"$set": toBson(request.Fields)}) != nil {
		panic("更新失败")
	}
	return &todo.EmptyResponse{}, nil
}

func toBson(fields map[string]string) bson.M {
	res := bson.M{}

	for k, v := range fields {
		if k == "status" {
			var err error
			res[k], err = strconv.ParseInt(v, 10, 32)
			if err != nil {
				panic(err)
			}
			res["doneAt"] = time.Now().UnixNano() / 1000
		} else {
			res[k] = v
		}
	}
	return res
}
