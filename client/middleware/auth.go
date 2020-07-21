package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"todomvc/client/core/codes"
	"todomvc/core/dao"
	"todomvc/core/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("正在访问路径", ctx.Request.URL.Path)
		requestUrl := ctx.Request.RequestURI
		if requestUrl != "/login" && requestUrl != "/register" {
			if utils.CheckJWTValid(ctx.GetHeader("accessToken")) {
				connection := dao.GetRedisConnection()
				if err := connection.Send("get", ctx.GetHeader("accessToken")); err != nil {
					panic(err)
				}
				if err := connection.Flush(); err != nil {
					panic(err)
				}
				if res, err := connection.Receive(); err != nil {
					panic(err)
				} else {
					if res == nil {
						ctx.String(codes.UNAUTHORIZED, "token 不受信任,请重新登录")
						ctx.Abort()
					} else {
						ctx.Next()
					}
				}
			} else {
				ctx.String(codes.UNAUTHORIZED, "token 过期，请重新登录")
				ctx.Abort()
			}
		} else {
			ctx.Next()
		}
		log.Println("访问结束")
		ctx.Abort()

	}
}
