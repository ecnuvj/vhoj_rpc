package rpc_remote

import (
	"github.com/ecnuvj/vhoj_rpc/model/remotepb"
	"google.golang.org/grpc"
	"log"
)

var RemoteServiceClient remotepb.RemoteServiceClient

func Init() {
	conn, err := grpc.Dial(":50005", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("remote rpc service connect...")
	RemoteServiceClient = remotepb.NewRemoteServiceClient(conn)
}
