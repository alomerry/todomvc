package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"todomvc/proto/user"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
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
		log.Fatal(err)
	}
	fmt.Println(reply)
}
