package service

import (
	"fmt"
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
	fmt.Printf("name:%s,password:%s,repeatPassword:%s\n", credential.Name, credential.Password, credential.RepeatPassword)
	if credential.Password != credential.RepeatPassword {
		return nil, status.New(codes.InvalidArgument, "两次密码不一致！").Err()
	}
	var result model.User
	if err := dao.DB.FindOne("user", bson.M{"name": credential.Name}, result); err == nil {
		return nil, status.New(codes.AlreadyExists, "用户名已存在！").Err()
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
