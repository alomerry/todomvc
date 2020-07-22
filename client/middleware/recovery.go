package middleware

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"todomvc/client/core/codes"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				ctx.String(codes.INTERNAL_SERVER_ERROR, "内部错误")
			}
		}()
		ctx.Next()
	}
}
