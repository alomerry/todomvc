package service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/service/user/model"
)
import "todomvc/proto/user"

func (*UserService) Login(ctx context.Context, userCredential *user.UserCredential) (*user.LoginResponse, error) {
	var result model.User
	if ok, err := isLoginValid(userCredential); !ok {
		return nil, err
	}
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name}, &result); err != nil {
		return nil, status.New(codes.InvalidArgument, "用户名不存在！").Err()
	}
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name, "password": userCredential.Password}, &result); err != nil {
		return nil, status.New(codes.PermissionDenied, "密码错误！").Err()
	}
	token := utils.MakeJWT(result.Id.Hex(), result.Name)

	con := dao.GetRedisConnection()
	if err := con.Send("setex", token, utils.GetTokenEcpirationTime(), result.Id.Hex()); err != nil {
		panic(err)
	}
	if err := con.Flush(); err != nil {
		panic(err)
	}
	if r, err := con.Receive(); err != nil {
		panic(err)
	} else {
		if r != nil {
			return &user.LoginResponse{
				Id:          result.Id.Hex(),
				Name:        result.Name,
				AccessToken: token,
			}, nil
		} else {
			panic("redis插入失败")
		}
	}
}

func isLoginValid(credential *user.UserCredential) (bool, error) {
	if len(credential.Name) < 6 || len(credential.Name) > 12 {
		return false, status.Error(codes.InvalidArgument, "用户名长度不合法")
	} else if len(credential.Password) < 6 || len(credential.Password) > 18 {
		return false, status.Error(codes.InvalidArgument, "密码长度不合法")
	}
	return true, nil
}
