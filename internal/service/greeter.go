package service

import (
	"context"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
)

// Greeter ...
type Greeter struct{}

// NewGreeterServer ...
func NewGreeterServer() *Greeter {
	return &Greeter{}
}

// SayHello ...
func (g *Greeter) SayHello(ctx context.Context, hw *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: hw.GetName()}, nil
}
