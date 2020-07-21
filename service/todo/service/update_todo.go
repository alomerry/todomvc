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
	err := dao.DB.UpdateById("todo", bson.ObjectIdHex(request.Id),
		bson.M{"$set": toBson(request.Fields)})
	if err != nil {
		panic(err)
	}
	return &todo.EmptyResponse{}, nil
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
			res["doneAt"] = time.Now().UnixNano() / 1000
		} else {
			res[k] = v
		}
	}
	return res
}
