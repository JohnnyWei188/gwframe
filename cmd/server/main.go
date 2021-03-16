package main

import (
	"github.com/JohnnyWei188/gwframe/internal/server"
	"github.com/JohnnyWei188/gwframe/internal/service"
	"github.com/JohnnyWei188/gwframe/pkg/gwframe"
	"github.com/JohnnyWei188/gwframe/pkg/transport/grpc"
	"github.com/JohnnyWei188/gwframe/pkg/transport/http"
)

func main() {
	app, err := initApp()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initApp() (*gwframe.App, error) {
	greeterService := service.NewGreeterServer()
	httpServer := server.NewHTTPServer(greeterService)
	grpcServer := server.NewGRPCServer(greeterService)
	app := newApp(httpServer, grpcServer)
	return app, nil
}

func newApp(hs *http.Server, gs *grpc.Server) *gwframe.App {
	return gwframe.New(
		gwframe.WithServer(
			hs,
			gs,
		),
	)
}
