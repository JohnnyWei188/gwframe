package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JohnnyWei188/gwframe/api/helloworld/v1"
	"github.com/JohnnyWei188/gwframe/internal/service"
	myhttp "github.com/JohnnyWei188/gwframe/internal/transport/http"
	"github.com/JohnnyWei188/gwframe/internal/transport/middleware"
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
	opts = append(opts, myhttp.HandleFunc("/test", MyTest()))
	opts = append(opts, myhttp.HandleFunc("/test2", MyTest()))
	opts = append(opts, myhttp.Middleware(middleware.Chain(middlewareHandler, middlewareHandler2)))

	mux := runtime.NewServeMux()
	helloworld.RegisterGreeterHandlerServer(context.Background(), mux, hw)
	srv := myhttp.NewServer(mux, opts...)
	return srv
}

// MyTest ...
func MyTest() http.HandlerFunc {
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
func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		cost := time.Now().Sub(start)
		log.Printf("path1111: %v, cost: %v \n", r.URL.Path, cost)
	})
}

func middlewareHandler2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		cost := time.Now().Sub(start)
		log.Printf("path2222: %v, cost: %v \n", r.URL.Path, cost)
	})
}

// swaggerServer returns swagger specification files located under "/swagger/"
//func swaggerServer(dir string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
//			http.NotFound(w, r)
//			return
//		}
//		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
//		p = path.Join(dir, p)
//		http.ServeFile(w, r, p)
//	}
//}
//
//// allowCORS allows Cross Origin Resoruce Sharing from any origin.
//func allowCORS(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if origin := r.Header.Get("Origin"); origin != "" {
//			w.Header().Set("Access-Control-Allow-Origin", origin)
//			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
//				preflightHandler(w, r)
//				return
//			}
//		}
//		h.ServeHTTP(w, r)
//	})
//}
//
//// preflightHandler adds the necessary headers in order to serve
//// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
//// We insist, don't do this without consideration in production systems.
//func preflightHandler(w http.ResponseWriter, r *http.Request) {
//	headers := []string{"Content-Type", "Accept", "Authorization"}
//	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
//	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
//	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
//}
//
//// healthzServer returns a simple health handler which returns ok.
//func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "text/plain")
//		if s := conn.GetState(); s != connectivity.Ready {
//			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
//			return
//		}
//		fmt.Fprintln(w, "ok")
//	}
//}
