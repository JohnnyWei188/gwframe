package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/JohnnyWei188/gwframe/internal/transport"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var _ transport.Server = (*Server)(nil)

// Server http server
type Server struct {
	*http.Server
	lis        net.Listener
	network    string
	address    string
	timeout    time.Duration
	mux        *http.ServeMux
	middleware Middleware
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

// WithMiddleware set middleware
func WithMiddleware(m Middleware) ServerOption {
	return func(s *Server) {
		s.middleware = m
	}
}

// HandleFunc set handle func, you can call this func many times with distinct pattern
func HandleFunc(pattern string, f http.HandlerFunc) ServerOption {
	return func(s *Server) {
		s.mux.HandleFunc(pattern, f)
	}
}

// NewServer create a server
func NewServer(mux *runtime.ServeMux, opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: time.Second,
		mux:     http.NewServeMux(),
	}
	for _, opt := range opts {
		opt(srv)
	}
	srv.mux.Handle("/", mux)
	srv.Server = &http.Server{
		Handler: srv.middleware(srv),
	}
	return srv
}

// ServeHTTP should write reply headers and data to the ResponseWriter
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), s.timeout)
	defer cancel()
	s.mux.ServeHTTP(res, req.WithContext(ctx))
}

// Start start
func (s *Server) Start() error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	if err := s.Serve(lis); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop stop
func (s *Server) Stop() error {
	return s.Shutdown(context.Background())
}

// Endpoint endpoints
func (s *Server) Endpoint() ([]string, error) {
	return []string{}, nil
}
