package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"todomvc/proto/user"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	loginClient := user.NewUserServiceClient(conn)
	reply, err := loginClient.Login(context.Background(), &user.UserCredential{
		Name:     "alomerry",
		Password: "120211",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
