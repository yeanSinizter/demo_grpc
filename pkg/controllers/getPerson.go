package controllers

import (
	"context"
	"demo_protoc/pkg/pb"
)

func (c *controller) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	return &pb.GetPersonResponse{
		Status: "Hello Person",
	}, nil
}
