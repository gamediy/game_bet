package tron

import (
	"bet/net/tron/pkg/client"
	"bet/net/tron/pkg/keystore"
	"fmt"
	"google.golang.org/grpc"
)

var (
	TronGrpcCLient = &client.GrpcClient{}
	err            error
)

func init() {
	keystore.ImportFromPrivateKey("1232")
	TronGrpcCLient, err = client.NewGrpcClient("grpc.trongrid.io:50051", grpc.WithInsecure())
	if err == nil {
		fmt.Println(err)
	}
}
