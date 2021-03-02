package rpc_submitter

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	"github.com/ecnuvj/vhoj_rpc/model/userpb"
	"google.golang.org/grpc"
	"log"
)

var UserServiceClient userpb.UserServiceClient

func Init() {
	conn, err := grpc.Dial(rpc_config.UserRpc.GetFullAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("user rpc service connect...")
	UserServiceClient = userpb.NewUserServiceClient(conn)
}
