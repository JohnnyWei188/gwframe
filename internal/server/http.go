package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
	"github.com/JohnnyWei188/gwframe/internal/service"
	myhttp "github.com/JohnnyWei188/gwframe/pkg/transport/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

// NewHTTPServer create a grpc server
func NewHTTPServer(hw *service.Greeter) *myhttp.Server {
	network := "tcp"
	addr := ":8019"
	timeout := time.Second

	var opts = []myhttp.ServerOption{}
	if network != "" {
		opts = append(opts, myhttp.Network(network))
	}
	if addr != "" {
		opts = append(opts, myhttp.Address(addr))
	}
	opts = append(opts, myhttp.Timeout(timeout))
	opts = append(opts, myhttp.HandleFunc("/test", CustomHandler()))
	opts = append(opts, myhttp.WithMiddleware(
		myhttp.Chain(
			CustomMiddleware, // 自定义middleware
		),
	))
	runtimeOpts := runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			OrigName:     true,
			EmitDefaults: true,
		},
	)
	mux := runtime.NewServeMux(runtimeOpts)
	helloworld.RegisterGreeterHandlerServer(context.Background(), mux, hw)
	srv := myhttp.NewServer(mux, opts...)
	return srv
}

// CustomHandler 自定义handler
func CustomHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ret := map[string]interface{}{
			"message": "测试看看",
		}
		b, _ := json.Marshal(ret)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

// CustomMiddleware 自定义middleware
func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		cost := time.Now().Sub(start)
		log.Printf("path: %v, cost: %v \n", r.URL.Path, cost)
	})
}
