package todo

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	todo "todomvc/client/model/todo"
	token "todomvc/core/utils"
	proto "todomvc/proto/todo"
)

func GetTodoServiceClient(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	var todoRequest todo.GetTodoRequest
	if ctx.ShouldBindQuery(&todoRequest) != nil {
		panic("参数绑定出错")
	}
	getTodoClient := proto.NewTodoServiceClient(connection)
	reply, err := getTodoClient.GetTodo(ctx, &proto.GetTodoRequest{
		SortBy:    todoRequest.SortBy,
		IsAscend:  todoRequest.IsAscend == 1,
		Page:      todoRequest.Page,
		PageSize:  todoRequest.PageSize,
		StartedAt: todoRequest.StartAt,
		EndedAt:   todoRequest.EndAt,
		Keyword:   todoRequest.Keyword,
		Status:    todoRequest.Status,
		UserId:    token.GetJWTClaims(ctx.GetHeader("accessToken"), "jti"),
	})
	if err != nil {
		panic(nil)
	}
	ctx.JSON(codes.OK, reply)
}
