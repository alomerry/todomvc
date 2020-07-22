package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	connector "todomvc/client/core/utils"
	todo "todomvc/client/model/todo"
	util "todomvc/core/utils"
	proto "todomvc/proto/todo"
)

func AddTodo(ctx *gin.Context) {
	connection := connector.GetConnection()
	defer connection.Close()

	var todo todo.AddTodoRequest
	if err := ctx.ShouldBind(&todo); err != nil {
		panic(err)
	}

	todoClient := proto.NewTodoServiceClient(connection)

	_, err := todoClient.AddTodo(context.Background(), &proto.AddTodoRequest{
		Color:   todo.Color,
		Content: todo.Content,
		UserId:  util.GetJWTClaims(ctx.GetHeader("accessToken"), "jti"),
	})
	if err != nil {
		panic(err)
	}

	ctx.String(codes.OK, "添加成功！")
}
