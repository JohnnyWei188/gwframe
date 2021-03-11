package grpc

import (
	"net"
	"time"

	"google.golang.org/grpc"
)

// Server http server
type Server struct {
	*grpc.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
}

// ServerOption server options
type ServerOption func(*Server)

// Network set network with server
func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address set address with server
func Address(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

// Timeout set timeout with server
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// NewServer create a server
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: time.Second,
	}
	for _, opt := range opts {
		opt(srv)
	}
	srv.Server = grpc.NewServer()
	return srv
}

// Start start
func (s *Server) Start() error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	return s.Serve(lis)
}

// Stop stop
func (s *Server) Stop() error {
	s.GracefulStop()
	return nil
}

// Endpoint endpoints
func (s *Server) Endpoint() ([]string, error) {
	return []string{}, nil
}
