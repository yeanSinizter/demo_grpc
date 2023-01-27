package controllers

import "demo_protoc/pkg/pb"

type controller struct {
	pb.UnsafePersonManagementServer
}

func NewController() *controller {
	return &controller{}
}
