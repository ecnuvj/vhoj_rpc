package client

import (
	"context"
	"encoding/json"
	"fmt"
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
		SourceCode: "#include <bits/stdc++.h>" +
			"using namespace std;" +
			"int main() {" +
			"int a,b;" +
			"while(cin >> a >> b) {" +
			"cout << a + b << endl;" +
			"}" +
			"return 0;" +
			"}",
		RemoteProblemId: "1000",
		RemoteOj:        "HDU",
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
