package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/user"
)

func RegisterServiceClient(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	registerClient := user.NewUserServiceClient(conn)

	name, password, repeatPassword := ctx.PostForm("name"), ctx.PostForm("password"), ctx.PostForm("repeatPassword")
	replay, err := registerClient.Register(context.Background(), &user.RegisterCredential{
		Name:           name,
		Password:       password,
		RepeatPassword: repeatPassword,
	})
	if err != nil {
		log.Print(err)
		ctx.String(codes.BAD_REQUEST, "用户已存在!")
		return
	} else {
		ctx.String(codes.OK, replay.AccessToken)
	}
}
