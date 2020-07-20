package utils

import (
	"google.golang.org/grpc"
	"sync"
	"todomvc/client/core/constant"
	"todomvc/core/utils"
)

type Conf struct {
	IpHost string
	Port   string
}

var (
	once sync.Once
	cfg  Conf
)

func GetDialHost() string {

	utils.Config(&once, &cfg, constant.ClientConf)
	ip, port := cfg.IpHost, cfg.Port
	return ip + ":" + port
}

func GetConnection() *grpc.ClientConn {
	connection, err := grpc.Dial(GetDialHost(), grpc.WithInsecure())
	if err != nil {
		panic(err)
		return nil
	}
	return connection
}
