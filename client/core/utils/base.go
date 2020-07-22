package utils

import (
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"todomvc/client/core/constant"
	"todomvc/core/utils"
)

type Conf struct {
	Host string
	Port string
}

var (
	once sync.Once
	cfg  Conf
)

func init() {
	utils.Config(&once, &cfg, constant.ClientConf)
}

func GetDialHost() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func GetConnection() *grpc.ClientConn {
	connection, err := grpc.Dial(GetDialHost(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return connection
}
