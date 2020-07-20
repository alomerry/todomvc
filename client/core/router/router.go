package router

import (
	"github.com/gin-gonic/gin"
	"todomvc/client/controller/todo"
	"todomvc/client/controller/user"
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
	router.POST("/login", user.LoginServiceClient)
	router.POST("/register", user.RegisterServiceClient)
	router.GET("/todos", todo.GetTodoServiceClient)
	router.POST("/todos", todo.AddTodoServiceClient)
	router.DELETE("/todo/:id", todo.RemoveTodoServiceClient)
	router.PUT("/todo/:id", todo.UpdateTodoServiceClient)
	//router.PUT("/test", gin.HandlerFunc(func(ctx *gin.Context) {
	//	fmt.Println(reflect.TypeOf(ctx.PostForm("status")))
	//	fmt.Println(ctx.PostForm("status"))
	//	fmt.Println(ctx.PostForm("content"))
	//	fmt.Println(ctx.PostForm("color"))
	//}))

}
