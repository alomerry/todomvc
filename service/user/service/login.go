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

func (*UserService) Login(ctx context.Context, userCredential *proto.UserCredential) (*proto.LoginResponse, error) {
	var user model.User
	if ok, err := isLoginValid(userCredential); !ok {
		return nil, err
	}
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name}, &user); err != nil {
		return nil, status.New(codes.InvalidArgument, "用户名不存在！").Err()
	}
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name, "password": userCredential.Password}, &user); err != nil {
		return nil, status.New(codes.PermissionDenied, "密码错误！").Err()
	}
	token := utils.MakeJWT(user.Id.Hex(), user.Name)
	_, err := dao.SetEx(token, utils.GetTokenEcpirationTime(), user.Id.Hex())
	if err != nil {
		panic(err)
	}
	return &proto.LoginResponse{
		Id:          user.Id.Hex(),
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

func isLoginValid(credential *proto.UserCredential) (bool, error) {
	if len(credential.Name) < 6 || len(credential.Name) > 12 {
		return false, status.Error(codes.InvalidArgument, "用户名长度不合法")
	} else if len(credential.Password) < 6 || len(credential.Password) > 18 {
		return false, status.Error(codes.InvalidArgument, "密码长度不合法")
	}
	return true, nil
}
