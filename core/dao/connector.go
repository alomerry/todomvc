package dao

import (
	"gopkg.in/mgo.v2"
	"sync"
	"todomvc/core/codes"
	"todomvc/core/utils"
)

type Conf struct {
	UserName string
	Password string
	IpHost   string
	Port     string
	DbName   string
}

var (
	cfg       Conf
	once      sync.Once
	dbSession *mgo.Session
)

func GetSession() *mgo.Session {
	return dbSession
}

func InitWithDBConf() (*mgo.Session, error) {
	utils.Config(&once, &cfg, codes.TodoConf)
	dialInfo, err := mgo.ParseURL(utils.GetConfItem(&cfg, "IpHost") + ":" + utils.GetConfItem(&cfg, "Port"))
	if err != nil {
		return nil, err
	}
	dialInfo.Database = utils.GetConfItem(&cfg, "DbName")
	dbSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return dbSession, err
}
