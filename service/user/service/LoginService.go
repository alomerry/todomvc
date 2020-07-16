package service

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"log"
	"todomvc/core/dao"
	"todomvc/core/utils"
	"todomvc/service/user/model"
)
import "todomvc/proto/user"

func (*UserService) Login(ctx context.Context, userCredential *user.UserCredential) (*user.LoginResponse, error) {
	var result model.User
	if err := dao.DB.FindOne("user", bson.M{"name": userCredential.Name, "password": userCredential.Password}, &result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	token := utils.MakeToken(result.Id.String(), result.Name)
	reply := &user.LoginResponse{
		Id:          result.Id.String(),
		Name:        result.Name,
		AccessToken: token,
	}
	return reply, nil
}
