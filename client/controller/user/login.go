package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"todomvc/client/core/codes"
	"todomvc/client/core/utils"
	"todomvc/client/model/user"
	proto "todomvc/proto/user"
)

func Login(ctx *gin.Context) {
	connection := utils.GetConnection()
	defer connection.Close()

	var user user.LoginRequest
	if ctx.ShouldBind(&user) != nil {
		panic("参数绑定错误")
	}

	userClient := proto.NewUserServiceClient(connection)

	reply, err := userClient.Login(context.Background(), &proto.UserCredential{
		Name:     user.Name,
		Password: user.Password,
	})
	if err != nil {
		log.Print(err)
		ctx.String(codes.NOT_FOUND, "用户名不存在!")
		return
	} else {
		ctx.JSON(codes.OK, reply)
	}
}
