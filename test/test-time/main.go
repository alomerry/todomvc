package main

import (
	"fmt"
	"time"
)

func main() {
	//
	timestamp := time.Now().UnixNano()
	fmt.Println(timestamp)
	timestamp = time.Now().Unix()
	fmt.Println(timestamp)
	fmt.Println("1594464695508\n")
	timestamp = 1594464695508
	fmt.Println(time.Unix(timestamp, 0))

}
