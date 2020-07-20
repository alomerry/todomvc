package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v = false
	test(&v)
	fmt.Println(v)
}

func test(v *bool) {
	*v, _ = strconv.ParseBool("true")
}
