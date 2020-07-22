package dao

import (
	"errors"
	"fmt"
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

func init() {
	utils.Config(&redisOnce, &redisConfig, constant.TodoConf)
	InitRedis()
}

func InitRedis() redis.Conn {
	connection, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.IpHost, redisConfig.RedisPort))
	if err != nil {
		panic(err)
	}
	return connection
}

func GetRedisConnection() redis.Conn {
	if connection == nil {
		return InitRedis()
	} else {
		return connection
	}
}

func Get(key string) (interface{}, error) {
	connection := GetRedisConnection()
	if err := connection.Send("get", key); err != nil {
		return nil, err
	}
	if err := connection.Flush(); err != nil {
		return nil, err
	}
	if result, err := connection.Receive(); err != nil {
		return nil, err
	} else if result == nil {
		return "", errors.New(fmt.Sprintf("key:[%s]不存在", key))
	} else {
		return result, nil
	}
}

func Set(key, value string) (interface{}, error) {
	connection := GetRedisConnection()
	if err := connection.Send("set", key, value); err != nil {
		return nil, err
	}
	if err := connection.Flush(); err != nil {
		return nil, err
	}
	if result, err := connection.Receive(); err != nil {
		return nil, err
	} else if result == nil {
		return nil, errors.New(fmt.Sprintf("key:[%s]插入失败", key))
	} else {
		return result, nil
	}
}

func SetEx(key string, expiration int, value string) (interface{}, error) {
	connection := GetRedisConnection()
	if err := connection.Send("setex", key, expiration, value); err != nil {
		return nil, err
	}
	if err := connection.Flush(); err != nil {
		return nil, err
	}
	if result, err := connection.Receive(); err != nil {
		return nil, err
	} else if result == nil {
		return nil, errors.New(fmt.Sprintf("key:[%s]插入失败", key))
	} else {
		return result, nil
	}
}
