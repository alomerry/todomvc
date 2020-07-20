package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
	"todomvc/core/dao"
	"todomvc/core/interceptor"
	"todomvc/proto/todo"
	"todomvc/proto/user"
	todoService "todomvc/service/todo/service"
	userService "todomvc/service/user/service"
)

func initInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.LoggerInterceptor))
}

func main() {

	server := grpc.NewServer(initInterceptor())

	user.RegisterUserServiceServer(server, new(userService.UserService))
	todo.RegisterTodoServiceServer(server, new(todoService.TodoService))

	if session, err := dao.InitAndAuthenticate(); err != nil {
		panic(err)
	} else {
		dao.InitRepository(session)
	}

	dao.InitRedis()

	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server.Serve(lis)
}
