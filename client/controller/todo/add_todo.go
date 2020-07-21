package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func AddTodoServiceClient(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	color, content, accessToken :=
		ctx.PostForm("color"), ctx.PostForm("content"), ctx.GetHeader("accessToken")
	addTodoClient := todo.NewTodoServiceClient(conn)
	reply, err := addTodoClient.AddTodo(context.Background(), &todo.AddTodoRequest{
		Color:       color,
		Content:     content,
		AccessToken: accessToken,
	})
	if err != nil {
		panic(err)
	}
	ctx.String(codes.OK, reply.Message)
}
