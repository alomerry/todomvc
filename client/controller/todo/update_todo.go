package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/client/model/todo"
	proto "todomvc/proto/todo"
)

func UpdateTodoServiceClient(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	todoClient := proto.NewTodoServiceClient(connection)
	_, err := todoClient.UpdateTodo(context.Background(), &proto.UpdateTodoRequest{
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
	var todo todo.UpdateTodoRequest
	if ctx.ShouldBind(&todo) != nil {
		panic("数据绑定错误")
	}

	if todo.Status == "1" {
		res["status"] = "true"
	}
	if todo.Status == "0" {
		res["status"] = "false"
	}
	if todo.Color != "" {
		res["color"] = todo.Color
	}
	if todo.Content != "" {
		res["content"] = todo.Content
	}
	return res
}
