package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"todomvc/client/core/utils"
	"todomvc/proto/user"
)

func RegisterController(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	registerClient := user.NewUserServiceClient(conn)

	name, password, repeatPassword := ctx.Request.Form.Get("name"), ctx.Request.Form.Get("password"), ctx.Request.Form.Get("repeatPassword")

	reply, err := registerClient.Register(context.Background(), &user.RegisterCredential{
		Name:           name,
		Password:       password,
		RepeatPassword: repeatPassword,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
