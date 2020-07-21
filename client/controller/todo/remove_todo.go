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
	id, accessToken :=
		ctx.Param("id"), ctx.GetHeader("accessToken")
	removeTodoClient := todo.NewTodoServiceClient(conn)
	reply, err := removeTodoClient.RemoveTodo(context.Background(), &todo.RemoveTodoRequest{
		Id:          id,
		AccessToken: accessToken,
	})
	if err != nil {
		panic(err)
	}
	ctx.String(codes.OK, reply.Message)
}
