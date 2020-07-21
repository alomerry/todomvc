package middleware

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/core/codes"
	"todomvc/core/dao"
	"todomvc/core/utils"
)

type Void struct {
}

// 无需验证 token 的路径
var paths map[string]Void
var void Void

func init() {
	paths = make(map[string]Void)
	paths["/login"] = void
	paths["/register"] = void
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, isExist := paths[ctx.Request.RequestURI]; isExist || ValidToken(ctx) {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}

func ValidToken(ctx *gin.Context) bool {
	_, err := dao.Get(ctx.GetHeader("accessToken"))
	switch {
	case !utils.IsJWTValid(ctx.GetHeader("accessToken")):
		ctx.String(codes.UNAUTHORIZED, "token 过期，请重新登录")
		return false
	case err != nil:
		ctx.String(codes.UNAUTHORIZED, "token 不受信任,请重新登录")
		return false
	default:
		return true
	}
}
