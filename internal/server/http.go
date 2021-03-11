package server

import (
	"time"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
	"github.com/JohnnyWei188/gwframe/internal/service"
	"github.com/JohnnyWei188/gwframe/internal/transport/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

// NewHTTPServer create a grpc server
func NewHTTPServer(hw *service.Greeter) *http.Server {
	network := "tcp"
	addr := ":8019"
	timeout := time.Second

	var opts = []http.ServerOption{}
	if network != "" {
		opts = append(opts, http.Network(network))
	}
	if addr != "" {
		opts = append(opts, http.Address(addr))
	}
	opts = append(opts, http.Timeout(timeout))

	mux := runtime.NewServeMux()
	helloworld.RegisterGreeterHandlerServer(context.Background(), mux, hw)
	srv := http.NewServer(mux, opts...)
	return srv
}
