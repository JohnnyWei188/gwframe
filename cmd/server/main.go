package main

import (
	"github.com/JohnnyWei188/gwframe/internal/server"
	"github.com/JohnnyWei188/gwframe/internal/service"
)

func main() {

	hw := service.NewGreeterServer()
	grpcSrv := server.NewGRPCServer(hw)
	go grpcSrv.Start()
	httpSrv := server.NewHTTPServer(hw)
	httpSrv.Start()
}
