package todo

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/core/codes"
	connector "todomvc/client/core/utils"
	todo "todomvc/client/model/todo"
	utils "todomvc/core/utils"
	proto "todomvc/proto/todo"
)

func GetTodo(ctx *gin.Context) {
	connection := connector.GetConnection()
	defer connection.Close()

	var todoRequest todo.GetTodoRequest
	if err := ctx.ShouldBindQuery(&todoRequest); err != nil {
		panic(err)
	}

	getTodoClient := proto.NewTodoServiceClient(connection)

	reply, err := getTodoClient.GetTodo(ctx, &proto.GetTodoRequest{
		SortBy:    todoRequest.SortBy,
		IsAscend:  todoRequest.IsAscend.Value,
		Page:      todoRequest.Page,
		PageSize:  todoRequest.PageSize,
		StartedAt: todoRequest.StartAt,
		EndedAt:   todoRequest.EndAt,
		Keyword:   todoRequest.Keyword,
		Status:    proto.Status(todoRequest.Status.Value),
		UserId:    utils.GetJWTClaims(ctx.GetHeader("accessToken"), "jti"),
	})
	if err != nil {
		panic(err)
	}

	ctx.JSON(codes.OK, reply)
}
