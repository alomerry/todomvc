package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func RemoveTodoServiceClient(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	todoClient := todo.NewTodoServiceClient(conn)
	_, err := todoClient.RemoveTodo(context.Background(), &todo.RemoveTodoRequest{
		Id: ctx.Param("id"),
	})
	if err != nil {
		panic(err)
	}
	ctx.String(codes.OK, "删除成功！")
}
