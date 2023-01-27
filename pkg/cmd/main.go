package main

import (
	"demo_protoc/pkg/controllers"
	"demo_protoc/pkg/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	Server := controllers.NewController()

	pb.RegisterPersonManagementServer(grpcServer, Server)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Fail to start service")
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Fail to Serve")
	}
}
