package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gopkg.in/mgo.v2"
	"sync"
	"time"
	"todomvc/core/codes"
)

type Conf struct {
	salt string
}

var (
	cfg       Conf
	once      sync.Once
	dbSession *mgo.Session
)

func MakeToken(id, name string) string {
	Config(&once, &cfg, codes.SaltConf)
	return md5V(id + (string(time.Now().Nanosecond())) + cfg.salt + name)
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
