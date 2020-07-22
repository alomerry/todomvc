package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"sync"
	"todomvc/core/constant"
	"todomvc/core/utils"
)

type Conf struct {
	UserName string
	Password string
	Host     string
	Port     string
	DbName   string
}

var (
	cfg       Conf
	once      sync.Once
	dbSession *mgo.Session
)

func init() {
	utils.Config(&once, &cfg, constant.TodoConf)
}

func GetSession() *mgo.Session {
	return dbSession
}

func InitAndAuthenticate() (*mgo.Session, error) {
	return initConnection("admin")
}

func initConnection(dbName string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
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
