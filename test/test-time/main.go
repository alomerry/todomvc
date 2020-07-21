package main

import (
	"fmt"
	"time"
)

func main() {
	//
	timestamp := time.Now().UnixNano()/1000
	fmt.Println(timestamp)
	timestamp = time.Now().Unix()
	fmt.Println(timestamp)
	fmt.Println("1595295812504")
	timestamp = 1595295812504
	fmt.Println(time.Unix(timestamp/1000,0))

}
