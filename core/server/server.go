package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net"
	"todomvc/core/dao"
	"todomvc/core/interceptor"
	"todomvc/proto/user"
	"todomvc/service/user/model"
	"todomvc/service/user/service"
)

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Filter))
	user.RegisterUserServiceServer(server, new(service.UserService))

	session, err := dao.InitWithDBConf()
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	mongoDBRepository := dao.InitRepository(session)
	mongoDBRepository.FindOne("user", bson.M{"name": "a"}, &user)
	fmt.Println(user)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
