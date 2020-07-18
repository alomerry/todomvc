package service

import (
	"fmt"
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
	// todo 验证数据长度等等

	var result model.User
	fmt.Printf("name:%s,password:%s", userCredential.Name, userCredential.Password)
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name}, &result); err != nil {
		return nil, status.New(codes.InvalidArgument, "用户名不存在！").Err()
	}
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name, "password": userCredential.Password}, &result); err != nil {
		return nil, status.New(codes.PermissionDenied, "密码错误！").Err()
	}
	token := utils.MakeJWT(result.Id.Hex(), result.Name)

	con := dao.GetRedisConnection()
	if err := con.Send("setex", result.Id.Hex(), "600", token); err != nil {
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
