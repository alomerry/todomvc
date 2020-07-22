package todo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	todo "todomvc/client/model/todo"
	token "todomvc/core/utils"
	proto "todomvc/proto/todo"
)

func AddTodo(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	var todo todo.AddTodoRequest
	if ctx.ShouldBind(&todo) != nil {
		panic("参数绑定出错")
	}

	todoClient := proto.NewTodoServiceClient(connection)

	_, err := todoClient.AddTodo(context.Background(), &proto.AddTodoRequest{
		Color:   todo.Color,
		Content: todo.Content,
		UserId:  token.GetJWTClaims(ctx.GetHeader("accessToken"), "jti"),
	})
	if err != nil {
		panic(err)
	}

	ctx.String(codes.OK, "添加成功！")
}
