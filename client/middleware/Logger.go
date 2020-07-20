package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"todomvc/client/core/codes"
	"todomvc/core/utils"
)

func Logger() gin.HandlerFunc {
	//todo 验证token redis是否有？是否过期
	return func(ctx *gin.Context) {
		log.Println("正在访问路径", ctx.Request.URL.Path)
		requestUrl := ctx.Request.RequestURI
		if requestUrl != "/login" || requestUrl != "/register" {
			if utils.CheckJWTValid(ctx.GetHeader("accessToken")) {
				//todd 验证redis
				ctx.Next()
			} else {
				ctx.String(codes.UNAUTHORIZED, "token 过期，请重新登录")
			}
		} else {
			ctx.Next()
		}
		log.Println("访问结束")
	}
}
