package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/user"
)

func LoginController(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	loginClient := user.NewUserServiceClient(conn)
	name, password := ctx.PostForm("name"), ctx.PostForm("password")

	reply, err := loginClient.Login(context.Background(), &user.UserCredential{
		Name:     name,
		Password: password,
	})
	if err != nil {
		log.Print(err)
		ctx.String(codes.NOT_FOUND, "用户名不存在!")
		return
	}
	ctx.JSON(codes.OK, gin.H{
		"id":    reply.Id,
		"name":  reply.Name,
		"token": reply.AccessToken,
	})
}
