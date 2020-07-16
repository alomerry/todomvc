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

	session, err := dao.InitAndAuthenticate()
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	var users []model.User
	dao.InitWithDBConf()
	mongoDBRepository := dao.InitRepository(session)
	session.DB("todo").C("user").Find(nil).All(&users);
	mongoDBRepository.FindOne("user", bson.M{"name": "alomerry"}, &user)
	fmt.Println(user)
	fmt.Println(users)

	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
