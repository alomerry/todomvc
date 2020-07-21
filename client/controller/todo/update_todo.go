package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/todo"
)

func UpdateTodoServiceClient(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	todoClient := todo.NewTodoServiceClient(connection)
	_, err := todoClient.UpdateTodo(context.Background(), &todo.UpdateTodoRequest{
		Id:     ctx.Param("id"),
		Fields: makeUpdateField(ctx),
	})
	if err != nil {
		panic(err)
	}
	ctx.String(codes.OK, "修改成功！")
}

func makeUpdateField(ctx *gin.Context) map[string]string {
	res := make(map[string]string)
	status, color, content := ctx.PostForm("status"), ctx.PostForm("color"), ctx.PostForm("content")
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
