package rpc_crawler

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	"github.com/ecnuvj/vhoj_rpc/model/crawlerpb"
	"github.com/ecnuvj/vhoj_rpc/model/problempb"
	"google.golang.org/grpc"
	"log"
)

var CrawlerServiceClient crawlerpb.ProblemServiceClient

func Init() {
	conn, err := grpc.Dial(rpc_config.ProblemRpc.GetFullAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("problem rpc service connect...")
	ProblemServiceClient = problempb.NewProblemServiceClient(conn)
}
