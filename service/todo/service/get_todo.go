package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	proto "todomvc/proto/todo"
	"todomvc/service/todo/model"
)

func (*TodoService) GetTodo(ctx context.Context, todoRequest *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	var result []model.Todo

	selector := makeSelector(todoRequest)
	if err := dao.DB.FindWithSorterAndLimit("todo", bson.M{"$and": selector}, todoRequest.SortBy, int(todoRequest.Page), int(todoRequest.PageSize), &result); err != nil {
		panic(err)
	}
	total, err := dao.DB.GetCountWithSorter("todo", bson.M{"$and": selector}, todoRequest.SortBy)
	if err != nil {
		panic(err)
	}
	return &proto.GetTodoResponse{
		Todos: conventTodos(result),
		Total: int32(total),
	}, nil
}
func makeSelector(todoRequest *proto.GetTodoRequest) []bson.M {
	var selector []bson.M
	selector = append(selector, bson.M{"userId": bson.ObjectIdHex(todoRequest.UserId)})
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
	if todoRequest.Status != proto.Status_UNSET {
		selector = append(selector, bson.M{"status": todoRequest.Status})
	}
	if todoRequest.Keyword != "" {
		selector = append(selector, bson.M{"content": bson.M{
			"$regex": todoRequest.Keyword,
		}})
	}
	return selector
}
func conventTodos(todos []model.Todo) []*proto.Todo {
	result := make([]*proto.Todo, len(todos))
	for i, todo := range todos {
		result[i] = &proto.Todo{
			Id:        todo.Id.Hex(),
			UserId:    todo.UserId.Hex(),
			CreatedAt: todo.CreatedAt,
			DoneAt:    todo.DoneAt,
			Color:     todo.Color,
			Status:    proto.Status(todo.Status),
			Content:   todo.Content,
		}
	}
	return result
}
