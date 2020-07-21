package dao

import (
	"gopkg.in/mgo.v2"
	"sync"
	"todomvc/core/constant"
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

func InitAndAuthenticate() (*mgo.Session, error) {
	return initConnection("admin")
}

func initConnection(dbName string) (*mgo.Session, error) {
	utils.Config(&once, &cfg, constant.TodoConf)
	dialInfo, err := mgo.ParseURL(cfg.IpHost + ":" + cfg.Port)
	if err != nil {
		return nil, err
	}
	dialInfo.Database = dbName
	dialInfo.Username = cfg.UserName
	dialInfo.Password = cfg.Password
	dbSession, err := mgo.DialWithInfo(dialInfo)
	dbSession.Login(&mgo.Credential{Username: cfg.UserName, Password: cfg.Password})
	if err != nil {
		return nil, err
	}
	dbSession = dbSession.DB(cfg.DbName).Session
	return dbSession, err
}
