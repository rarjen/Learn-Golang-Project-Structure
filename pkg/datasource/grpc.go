package datasource

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpc(host string) (*grpc.ClientConn, error) {
	return grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
