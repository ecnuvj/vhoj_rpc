module github.com/ecnuvj/vhoj_rpc

go 1.14

require (
	github.com/ecnuvj/vhoj_common v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/ecnuvj/vhoj_common => ../vhoj_common
