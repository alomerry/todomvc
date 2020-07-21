package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/proto/todo"
	"todomvc/service/todo/model"
)

func (*TodoService) GetTodo(ctx context.Context, todoRequest *todo.GetTodoRequest) (*todo.GetTodoResponse, error) {
	var result []model.Todo
	var total int
	var err error
	userId := bson.ObjectIdHex(utils.GetJWTClaims(todoRequest.AccessToken, "jti"))
	var selector []bson.M
	selector = append(selector, bson.M{"belongTo": userId})
	if todoRequest.SortBy == "" {
		todoRequest.SortBy = "createdAt"
	}
	if todoRequest.StartedAt != 0 && todoRequest.EndedAt != 0 {
		selector = append(selector, bson.M{
			"createdAt": bson.M{
				"$lte": todoRequest.EndedAt * 1000,
				"$gte": todoRequest.StartedAt * 1000,
			},
		})
	}
	if todoRequest.Status != "" {
		tmp, err := strconv.ParseBool(todoRequest.Status)
		if err != nil {
			panic(err)
		}
		selector = append(selector, bson.M{"status": tmp})
	}
	if todoRequest.KeyWord != "" {
		selector = append(selector, bson.M{"content": bson.M{
			"$regex": todoRequest.KeyWord,
		}})
	}
	if err := dao.DB.FindWithSorterAndLimit("todo", bson.M{"$and": selector}, todoRequest.SortBy, int(todoRequest.Page), int(todoRequest.PageSize), &result); err != nil {
		panic(err)
	}
	total, err = dao.DB.GetCountWithSorter("todo", bson.M{"$and": selector}, todoRequest.SortBy)
	if err != nil {
		panic(err)
	}
	return &todo.GetTodoResponse{
		Todos: conventTodos(result),
		Total: int32(total),
	}, nil
}

func conventTodos(todos []model.Todo) []*todo.Todo {
	var res []*todo.Todo
	for _, item := range todos {
		res = append(res, &todo.Todo{
			Id:        item.Id.Hex(),
			BelongTo:  item.BelongTo.Hex(),
			CreatedAt: item.CreatedAt,
			DoneAt:    item.DoneAt,
			Color:     item.Color,
			Status:    item.Status,
			Content:   item.Content,
		})
	}
	return res
}
