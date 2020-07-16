package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(md5V("wjc1231231"))
}
