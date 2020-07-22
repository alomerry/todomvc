package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func RemoveTodo(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	todoClient := todo.NewTodoServiceClient(connection)

	_, err := todoClient.RemoveTodo(context.Background(), &todo.RemoveTodoRequest{
		Id: ctx.Param("id"),
	})
	if err != nil {
		panic(err)
	}

	ctx.String(codes.OK, "删除成功！")
}
