package interceptor

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"todomvc/core/utils"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("[info:%s]\n", info.FullMethod)
	//todo need delete
	utils.Log("info", "×××××", info.FullMethod, nil)
	resp, err := handler(ctx, req)
	if err != nil {
		utils.Log("error", "******", info.FullMethod, err)
	}
	return resp, err
}
