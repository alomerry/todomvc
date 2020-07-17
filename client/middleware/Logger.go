package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("正在访问路径",ctx.Request.URL.Path)
		ctx.Next()
		log.Println("访问结束")
	}
}