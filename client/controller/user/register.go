package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/client/model/user"
	proto "todomvc/proto/user"
)

func Register(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	var user user.RegisterRequest
	if err := ctx.ShouldBind(&user); err != nil {
		panic(err)
	}

	userClient := proto.NewUserServiceClient(connection)
	replay, err := userClient.Register(context.Background(), &proto.RegisterCredential{
		Name:           user.Name,
		Password:       user.Password,
		RepeatPassword: user.RepeatPassword,
	})
	if err != nil {
		log.Println(err)
		ctx.String(codes.BAD_REQUEST, "用户已存在!")
		return
	}
	ctx.String(codes.OK, replay.AccessToken)
}
