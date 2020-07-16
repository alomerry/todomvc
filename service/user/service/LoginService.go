package service

import "golang.org/x/net/context"
import "todomvc/proto/user"

func (*UserService) Login(ctx context.Context, userCredential *user.UserCredential) (*user.LoginResponse, error) {
	//username, password := userCredential.GetUsername(), userCredential.GetPassword()
	return nil, nil
}
