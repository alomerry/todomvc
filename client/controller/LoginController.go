package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/proto/user"

)

func LoginController(ctx *gin.Context) {
	conn := utils.GetConnection()
	defer conn.Close()
	loginClient := user.NewUserServiceClient(conn)
	name, password := ctx.PostForm("name"), ctx.PostForm("password")
	fmt.Printf("name:%s,password:%s\n", name, password)
	reply, err := loginClient.Login(context.Background(), &user.UserCredential{
		Name:     name,
		Password: password,
	})
	if err != nil {

		log.Println(reflect.TypeOf(err))
		ctx.String(codes.NOT_FOUND, err.Error())
		ctx.String(codes.UNAUTHORIZED, err.Error())
	}
	ctx.JSON(codes.OK, gin.H{
		"id":    reply.Id,
		"name":  reply.Name,
		"token": reply.AccessToken,
	})
}
