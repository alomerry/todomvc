package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"todomvc/client/core/codes"
	"todomvc/core/dao"
	"todomvc/core/utils"
)

func Authenticate() gin.HandlerFunc {
	//todo 验证token redis是否有？是否过期
	return func(ctx *gin.Context) {
		log.Println("正在访问路径", ctx.Request.URL.Path)
		requestUrl := ctx.Request.RequestURI
		if requestUrl != "/login" || requestUrl != "/register" {
			if utils.CheckJWTValid(ctx.GetHeader("accessToken")) {
				//todd 验证redis
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
						ctx.String(codes.UNAUTHORIZED, "token不受信任,请重新登录")
					} else {
						ctx.Next()
					}
				}
			} else {
				ctx.String(codes.UNAUTHORIZED, "token 过期，请重新登录")
			}
		} else {
			ctx.Next()
		}
		log.Println("访问结束")
	}
}
