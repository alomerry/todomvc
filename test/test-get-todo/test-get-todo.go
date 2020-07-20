package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"todomvc/proto/todo"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	getTodoClient := todo.NewTodoServiceClient(conn)
	reply, err := getTodoClient.GetTodo(context.Background(), &todo.GetTodoRequest{
		SortBy:      "",
		IsAscend:    false,
		Page:        1,
		PageSize:    10,
		AccessToken: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
