package dao

import (
	"fmt"
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
	return initConection(utils.GetConfItem(&cfg, "DbName"))
}

func InitAndAuthenticate() (*mgo.Session, error) {
	return initConection("admin")
}

func initConection(dbName string) (*mgo.Session, error) {
	utils.Config(&once, &cfg, codes.TodoConf)
	dialInfo, err := mgo.ParseURL(utils.GetConfItem(&cfg, "IpHost") + ":" + utils.GetConfItem(&cfg, "Port"))
	if err != nil {
		return nil, err
	}
	dialInfo.Database = dbName
	dialInfo.Username = utils.GetConfItem(&cfg, "UserName");
	dialInfo.Password = utils.GetConfItem(&cfg, "Password");
	fmt.Printf("%s,%s,%s",dialInfo.Username,dialInfo.Password)
	dbSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return dbSession, err
}