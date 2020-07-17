package interceptor

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}
