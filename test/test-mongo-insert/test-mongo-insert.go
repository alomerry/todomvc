package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	todo "todomvc/service/todo/model"
	user "todomvc/service/user/model"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	fmt.Println(err)
	var user user.User
	var todos []todo.Todo
	session.Login(&mgo.Credential{Username: "alomerry", Password: "120211"})
	err = session.DB("todo").C("user").Find(bson.M{"name": "alomerry"}).One(&user)
	//err = session.DB("todo").C("todo").Insert(bson.M{"belongTo": user.Id, "color": "#fff", "status": false, "content": "学习", "createdAt": time.Now().Unix()})
	err = session.DB("todo").C("todo").Find(nil).All(&todos)
	fmt.Println(err)
	fmt.Println(user)
	fmt.Println(todos)
	session.Close()
}
