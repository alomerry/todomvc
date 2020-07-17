package dao

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"todomvc/core/constant"
	"todomvc/core/utils"
)

type RedisConf struct {
	UserName  string
	Password  string
	IpHost    string
	Port      string
	DbName    string
	RedisPort string
}

var (
	redisConfig RedisConf
	redisOnce   sync.Once
)
var connection redis.Conn

func InitRedis() redis.Conn {
	utils.Config(&redisOnce, &redisConfig, constant.TodoConf)
	Connection, err := redis.Dial("tcp", redisConfig.IpHost+":"+redisConfig.RedisPort)
	if err != nil {
		panic(err)
	}
	return Connection
}

func GetRedisConnection() redis.Conn {
	if connection == nil {
		return InitRedis()
	} else {
		return connection
	}
}
