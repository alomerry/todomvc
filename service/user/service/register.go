package service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/proto/user"
	"todomvc/service/user/model"
)

func (*UserService) Register(ctx context.Context, credential *user.RegisterCredential) (*user.RegisterResponse, error) {
	if ok, err := isRegisterValid(credential); !ok {
		return nil, err
	}
	var result model.User
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name}, result); err == nil {
		return nil, status.Error(codes.AlreadyExists, "用户名已存在！")
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

func isRegisterValid(credential *user.RegisterCredential) (bool, error) {
	if len(credential.Name) < 6 || len(credential.Name) > 12 {
		return false, status.Error(codes.InvalidArgument, "用户名长度不合法")
	} else if len(credential.Password) < 6 || len(credential.Password) > 18 {
		return false, status.Error(codes.InvalidArgument, "密码长度不合法")
	} else
	if credential.Password != credential.RepeatPassword {
		return false, status.Error(codes.InvalidArgument, "两次密码不一致！")
	}
	return true, nil
}
