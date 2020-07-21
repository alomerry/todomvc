package interceptor

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("[info:%s]\n", info.FullMethod)
	resp, err := handler(ctx, req)
	fmt.Printf("[info:%s]\n", err)
	return resp, err
}
