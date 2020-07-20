package interceptor

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("check auth:根据包判断需不需要有token验证，进redis检查是否存在token，检验token是否合法")
	return handler(ctx, req)
}
