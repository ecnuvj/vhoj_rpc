package rpc_submitter

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	"github.com/ecnuvj/vhoj_rpc/model/submitterpb"
	"google.golang.org/grpc"
	"log"
)

var SubmitServiceClient submitterpb.SubmitServiceClient

func Init() {
	conn, err := grpc.Dial(rpc_config.SubmitterRpc.GetFullAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("submitter rpc service connect...")
	SubmitServiceClient = submitterpb.NewSubmitServiceClient(conn)
}
