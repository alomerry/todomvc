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

func (repo *mongoDBRepository) Find(collectionName string, selector bson.M, result interface{}) error {
	return repo.db.C(collectionName).Find(selector).All(result)
}

func (repo *mongoDBRepository) FindWithLimit(collectionName string, selector bson.M, page, pageSize int, result interface{}) error {
	return repo.db.C(collectionName).Find(selector).Skip((page - 1) * pageSize).Limit(pageSize).All(result)
}

func (repo *mongoDBRepository) GetCount(collectionName string, selector bson.M) (int, error) {
	return repo.db.C(collectionName).Find(selector).Count()
}

func (repo *mongoDBRepository) GetCountWithSorter(collectionName string, selector bson.M, field string) (int, error) {
	return repo.db.C(collectionName).Find(selector).Sort(field).Count()
}

func (repo *mongoDBRepository) FindWithSorterAndLimit(collectionName string, selector bson.M, field string, page, pageSize int, result interface{}) error {
	return repo.db.C(collectionName).Find(selector).Sort(field).Skip((page - 1) * pageSize).Limit(pageSize).All(result)
}

func (repo *mongoDBRepository) InsertOne(collectionName string, selector bson.M) error {
	return repo.db.C(collectionName).Insert(&selector)
}

func (repo *mongoDBRepository) RemoveId(collectionName string, id bson.ObjectId) error {
	return repo.db.C(collectionName).RemoveId(id)
}

func (repo *mongoDBRepository) UpdateId(collectionName string, id bson.ObjectId, update interface{}) error {
	return repo.db.C(collectionName).UpdateId(id, update)
}
