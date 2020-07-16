package interceptor

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func Filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("fileter:", info)
	return handler(ctx, req)
}
