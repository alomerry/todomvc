package service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/core/utils"
	proto "todomvc/proto/user"
	"todomvc/service/user/model"
)

func (*UserService) Register(ctx context.Context, credential *proto.RegisterCredential) (*proto.RegisterResponse, error) {
	if ok, err := isRegisterValid(credential); !ok {
		return nil, err
	}
	var user model.User
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name}, user); err == nil {
		return nil, status.Error(codes.AlreadyExists, "用户名已存在！")
	}
	if err := dao.DB.InsertOne("user", bson.M{"name": credential.Name, "password": credential.Password}); err != nil {
		panic(err)
	}
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name, "password": credential.Password}, &user); err != nil {
		panic(err)
	}
	return &proto.RegisterResponse{
		Id:          user.Id.Hex(),
		AccessToken: utils.MakeJWT(user.Id.Hex(), user.Name),
	}, nil
}

func isRegisterValid(credential *proto.RegisterCredential) (bool, error) {
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
