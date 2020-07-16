package dao

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseRepository interface {
	FindOne(collectionName string, selector bson.M, result interface{}) error
	FindOneWithSortor(ctx context.Context, collectionName string, selector bson.M, sortor []string, result interface{}) error
	FindAll(ctx context.Context, collectionName string, selector bson.M, sortor []string, limit int, result interface{}) error
	UpdateOne(ctx context.Context, collectionName string, selector bson.M, updator bson.M) error
	UpdateAll(ctx context.Context, collectionName string, selector bson.M, updator bson.M) (int, error)
	Insert(ctx context.Context, collectionName string, docs ...interface{}) error
	RemoveOne(ctx context.Context, collectionName string, selector bson.M) error
	RemoveAll(ctx context.Context, collectionName string, selector bson.M) (int, error)
	Count(ctx context.Context, collectionName string, selector bson.M) (int, error)
}

type mongoDBRepository struct {
	session *mgo.Session
	db      *mgo.Database
}

var DB *mongoDBRepository = &mongoDBRepository{
}

func InitRepository(session *mgo.Session) *mongoDBRepository {
	DB.session = session
	DB.db = session.DB("todo")
	return DB
}

func (repo *mongoDBRepository) FindOne(collectionName string, selector bson.M, result interface{}) error {
	return repo.db.C(collectionName).Find(selector).One(result)
}

func (repo *mongoDBRepository) InsertOne(collectionName string, selector bson.M) error {
	return repo.db.C(collectionName).Insert(&selector)
}
