package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/core/errors"
	"todomvc/core/utils"
	"todomvc/proto/user"
	"todomvc/service/user/codes"
	"todomvc/service/user/model"
)

func (*UserService) Register(ctx context.Context, credential *user.RegisterCredential) (*user.RegisterResponse, error) {
	if credential.Password != credential.RepeatPassword {
		return nil, errors.RPCError{Code: codes.TwoPasswordMismatch, Desc: "两次密码不一致！"}
	}
	var result model.User
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name}, result); err != nil {
		return nil, errors.RPCError{Code: codes.UserNameExist, Desc: "用户名已存在！"}
	}
	if err := dao.DB.InsertOne("user", bson.M{"name": credential.Name, "password": credential.Password}); err != nil {
		panic(err)
		return nil, err
	}
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name, "password": credential.Password}, &result); err != nil {
		panic(err)
		return nil, err
	}

	return &user.RegisterResponse{
		Id:          result.Id.Hex(),
		AccessToken: utils.MakeJWT(result.Id.Hex(), result.Name),
	}, nil
}
