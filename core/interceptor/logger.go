package interceptor

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("[info][%s]\n", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Printf("[%s][%s]\n", resp, err)
	return resp, err
}
