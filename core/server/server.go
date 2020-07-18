package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
	"todomvc/core/dao"
	"todomvc/core/interceptor"
	"todomvc/proto/user"
	"todomvc/service/user/service"
)

func initInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.Authenticate,
		interceptor.LoggerInterceptor))
}

func main() {

	server := grpc.NewServer(initInterceptor())

	user.RegisterUserServiceServer(server, new(service.UserService))

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
