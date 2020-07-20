package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func UpdateTodoServiceClient(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	id, accessToken := ctx.Param("id"), ctx.GetHeader("accessToken")

	updateTodoClient := todo.NewTodoServiceClient(conn)
	replay, err := updateTodoClient.UpdateTodo(context.Background(), &todo.UpdateTodoRequest{
		Id:          id,
		AccessToken: accessToken,
		Fields:      makeUpdateField(ctx),
	})
	if err != nil {
		panic(err)
	}
	ctx.String(codes.OK, replay.Message)
}
func makeUpdateField(ctx *gin.Context) map[string]string {
	res := make(map[string]string)
	status, color, content := ctx.PostForm("status"), ctx.PostForm("color"), ctx.PostForm("content")
	fmt.Println(status)
	fmt.Println(color)
	fmt.Println(content)
	if status != "" {
		res["status"] = status
	}
	if color != "" {
		res["color"] = color
	}
	if content != "" {
		res["content"] = content
	}

	return res
}
