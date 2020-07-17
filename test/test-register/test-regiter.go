package main

import (
	"context"
	"fmt"
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
	registerClient := user.NewUserServiceClient(conn)
	reply, err := registerClient.Register(context.Background(), &user.RegisterCredential{
		Name:           "alomerry",
		Password:       "120211",
		RepeatPassword: "120211",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
