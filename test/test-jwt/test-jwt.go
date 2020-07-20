package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"todomvc/core/utils"
)

func main() {
	token := utils.MakeJWT("1", "alomerry")
	fmt.Println(token)
	time.Sleep(time.Second * 2)
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Fatal(ok)
			return false, nil
		}
		return []byte("dasfdfd43454dE#%#"), nil
	})
	fmt.Println(err)
	var maps = jwtToken.Claims.(jwt.MapClaims)
	fmt.Println(maps["aud"])
}
