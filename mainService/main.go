package main

import (
	_ "github.com/hcdast/go-micro-frame/mainService/conf"
	"github.com/hcdast/go-micro-frame/mainService/handler"
	pb "github.com/hcdast/go-micro-frame/mainService/proto"
	"github.com/micro/micro/v3/service"

	"github.com/micro/micro/v3/service/logger"
)

const (
	ServerName = "go.micro.main.service" // server name
)

func main() {
	// Create service
	srv := service.New(
		service.Name(ServerName),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMainServiceHandler(srv.Server(), new(handler.MainService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
