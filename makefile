grpc: 
	protoc --go_out=./pkg/ --go-grpc_out=./pkg pkg/pb_gen/*.proto
	protoc --go_out=./../gateway --go-grpc_out=./../gateway pkg/pb_gen/*.proto
	