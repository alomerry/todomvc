package main

import (
	"fmt"
	"runtime"
)

func test1() {
	log()
	fmt.Println("test1 run.")
}

func log() {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		pcName := runtime.FuncForPC(pc).Name()
		fmt.Println(pcName)
	}
}

func main() {
	test1()
}
