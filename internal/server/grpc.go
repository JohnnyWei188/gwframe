package server

import (
	"time"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
	"github.com/JohnnyWei188/gwframe/internal/service"
	"github.com/JohnnyWei188/gwframe/internal/transport/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCServer create a grpc server
func NewGRPCServer(hw *service.Greeter) *grpc.Server {
	network := "tcp"
	addr := ":8018"
	timeout := time.Second

	var opts = []grpc.ServerOption{}
	if network != "" {
		opts = append(opts, grpc.Network(network))
	}
	if addr != "" {
		opts = append(opts, grpc.Address(addr))
	}
	opts = append(opts, grpc.Timeout(timeout))
	srv := grpc.NewServer(opts...)
	reflection.Register(srv.Server)
	helloworld.RegisterGreeterServer(srv.Server, hw)
	return srv
}
