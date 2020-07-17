package router

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/controller"
	"todomvc/client/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	initMiddleware(router)
	initController(router)
	return router
}

func initMiddleware(router *gin.Engine) {
	router.Use(middleware.Logger())
}

func initController(router *gin.Engine) {
	router.POST("/login", controller.LoginController)
	router.POST("/register", controller.RegisterController)
}
