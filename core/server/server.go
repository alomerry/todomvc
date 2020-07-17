package main

import (
	"google.golang.org/grpc"
	"net"
	"todomvc/core/dao"
	"todomvc/core/interceptor"
	"todomvc/proto/user"
	"todomvc/service/user/service"
)

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Authenticate))

	user.RegisterUserServiceServer(server, new(service.UserService))

	session, err := dao.InitAndAuthenticate()
	dao.InitRepository(session)

	dao.InitRedis()

	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server.Serve(lis)
}
