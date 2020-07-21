package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"time"
	"todomvc/core/constant"
)

type Conf struct {
	ACCESS_TOKEN_SALT     string
	JWT_SECRETKEY         string
	TOKEN_ECPIRATION_TIME int
}

var (
	cfg       Conf
	once      sync.Once
	dbSession *mgo.Session
)

func init() {
	Config(&once, &cfg, constant.SaltConf)
}

func GetTokenEcpirationTime() int {
	res := GetConfItem(&cfg, "TOKEN_ECPIRATION_TIME").Int()
	return int(res)
}

func MakeToken(id, name string) string {
	return md5V(id + (string(time.Now().Nanosecond())) + cfg.ACCESS_TOKEN_SALT + name)
}

func MakeJWT(id, name string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(cfg.TOKEN_ECPIRATION_TIME)).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "alomerry",
		Id:        id,
		Audience:  name,
	}

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(cfg.JWT_SECRETKEY))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("make token:", tokenString)
	return tokenString
}

func GetJWTClaims(accessToken, clasimName string) string {
	jwtToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			panic(ok)
			return false, nil
		}
		return []byte(cfg.JWT_SECRETKEY), nil
	})
	if err != nil {
		panic(err)
	}
	return (jwtToken.Claims.(jwt.MapClaims))[clasimName].(string)
}

func IsJWTValid(accessToken string) bool {
	Config(&once, &cfg, constant.SaltConf)
	jwtToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			panic(ok)
			return false, nil
		}
		return []byte(cfg.JWT_SECRETKEY), nil
	})
	if err != nil {
		return false
	}
	return jwtToken.Valid
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
