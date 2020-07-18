package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"todomvc/service/user/model"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	fmt.Print(err)
	var user model.User
	session.Login(&mgo.Credential{Username: "alomerry", Password: "120211"})
	err = session.DB("todo").C("user").Find(bson.M{"name": "misaki"}).One(&user)
	fmt.Print(err)
	fmt.Print(user)
	session.Close()
}
