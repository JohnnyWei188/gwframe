package main

import (
	"github.com/JohnnyWei188/gwframe/internal/server"
	"github.com/JohnnyWei188/gwframe/internal/service"
)

func main() {

	hw := service.NewGreeterServer()
	httpSrv := server.NewHTTPServer(hw)
	go httpSrv.Start()
	grpcSrv := server.NewGRPCServer(hw)
	grpcSrv.Start()
}
