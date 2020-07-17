package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"reflect"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	//c.Send("setex", "test", "3", "h324210f9a")
	//c.Flush()
	//r, err := c.Receive()
	//fmt.Println(r, err)

	c.Send("get", "test")
	c.Flush()
	r, err := c.Receive()
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(r)
	c.Close()

}
