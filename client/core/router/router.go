package router

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/controller/todo"
	"todomvc/client/controller/user"
	"todomvc/client/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	initMiddleware(router)
	initController(router)
	return router
}

func initMiddleware(router *gin.Engine) {
	router.Use(middleware.Authenticate())
}

func initController(router *gin.Engine) {
	router.POST("/login", user.Login)
	router.POST("/register", user.Register)
	router.GET("/todos", todo.GetTodo)
	router.POST("/todos", todo.AddTodo)
	router.DELETE("/todo/:id", todo.RemoveTodo)
	router.PUT("/todo/:id", todo.UpdateTodo)
}
