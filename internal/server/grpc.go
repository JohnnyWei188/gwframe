package server

import (
	"time"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
	"github.com/JohnnyWei188/gwframe/internal/service"
	mygrpc "github.com/JohnnyWei188/gwframe/pkg/transport/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// NewGRPCServer create a grpc server
func NewGRPCServer(hw *service.Greeter) *mygrpc.Server {
	network := "tcp"
	addr := ":8018"
	timeout := time.Second

	var opts = []mygrpc.ServerOption{}
	if network != "" {
		opts = append(opts, mygrpc.Network(network))
	}
	if addr != "" {
		opts = append(opts, mygrpc.Address(addr))
	}
	opts = append(opts, mygrpc.Timeout(timeout))
	opts = append(opts, mygrpc.Options(
		grpc.ChainUnaryInterceptor(
			UnaryServerInterceptor(&DemoLimiter{}), // 添加自定义拦截器
		),
	))

	srv := mygrpc.NewServer(opts...)
	reflection.Register(srv.Server)
	helloworld.RegisterGreeterServer(srv.Server, hw)
	return srv
}

// DemoLimiter this is a demo for customize interceptor
type DemoLimiter struct{}

// Limit ...
func (dl *DemoLimiter) Limit() bool {
	return false
}

// Limiter defines the interface to perform request rate limiting.
// If Limit function return true, the request will be rejected.
// Otherwise, the request will pass.
type Limiter interface {
	Limit() bool
}

// UnaryServerInterceptor returns a new unary server interceptors that performs request rate limiting.
func UnaryServerInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if limiter.Limit() {
			return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new stream server interceptor that performs rate limiting on the request.
func StreamServerInterceptor(limiter Limiter) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if limiter.Limit() {
			return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
		}
		return handler(srv, stream)
	}
}
