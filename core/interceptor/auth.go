package interceptor

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("check auth")
	return handler(ctx, req)
}
