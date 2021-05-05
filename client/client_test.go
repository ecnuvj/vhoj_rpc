package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_rpc/client/rpc_remote"
	"github.com/ecnuvj/vhoj_rpc/model/remotepb"
	"testing"
	"time"
)

func TestRemoteClient(t *testing.T) {
	rpc_remote.Init()
	request := &remotepb.SubmitCodeRequest{
		UserId:   1212,
		Username: "bqx",
		SourceCode: "#include <bits/stdc++.h>\n" +
			"using namespace std;\n" +
			"int main() {\n" +
			"int a,b;\n" +
			"while(cin >> a >> b) {\n" +
			"cout << a + b << endl;\n" +
			"}\n" +
			"return 0;\n" +
			"}",
		RemoteProblemId: "1000",
		RemoteOj:        int32(remote_oj.HDU),
		ProblemId:       2,
		Language:        1,
	}
	resp, err := rpc_remote.RemoteServiceClient.SubmitCode(context.Background(), request)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(resp)
	fmt.Println(string(str))

	time.Sleep(3 * time.Second)
}

func TestSubmitCode(t *testing.T) {
	rpc_remote.Init()
}

func TestCrawlProblem(t *testing.T) {
	rpc_remote.Init()
	req := &remotepb.CrawlProblemRequest{
		RemoteOj:        int32(remote_oj.JSK),
		RemoteProblemId: "A1003",
		Enforce:         true,
	}
	resp, err := rpc_remote.RemoteServiceClient.CrawlProblem(context.Background(), req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(resp)
	fmt.Println(string(str))
}
