package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strconv"
	"todomvc/client/core/codes"
	connector "todomvc/client/core/utils"
	"todomvc/client/model/todo"
	proto "todomvc/proto/todo"
)

func UpdateTodo(ctx *gin.Context) {
	connection := connector.GetConnection()
	defer connection.Close()

	todoClient := proto.NewTodoServiceClient(connection)

	fields, err := makeUpdateField(ctx)
	if err != nil {
		panic(err)
	}
	_, err = todoClient.UpdateTodo(context.Background(), &proto.UpdateTodoRequest{
		Id:     ctx.Param("id"),
		Fields: fields,
	})
	if err != nil {
		panic(err)
	}

	ctx.String(codes.OK, "修改成功！")
}

func makeUpdateField(ctx *gin.Context) (map[string]string, error) {
	needUpdateFields := make(map[string]string)
	var todo todo.UpdateTodoRequest
	if err := ctx.ShouldBind(&todo); err != nil {
		return nil, err
	}
	fmt.Println(todo.Status.Value)
	switch {
	case todo.Status.Value == 1 || todo.Status.Value == 2:
		needUpdateFields["status"] = strconv.Itoa(todo.Status.Value)
	}
	if todo.Color != "" {
		needUpdateFields["color"] = todo.Color
	}
	if todo.Content != "" {
		needUpdateFields["content"] = todo.Content
	}
	return needUpdateFields, nil
}
