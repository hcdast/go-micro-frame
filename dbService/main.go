package main

import (
	"github.com/hcdast/go-micro-frame/dbService/handler"
	_ "github.com/hcdast/go-micro-frame/dbService/models"
	pb "github.com/hcdast/go-micro-frame/dbService/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

const (
	ServerName = "go.micro.db.service" // server name
)

func main() {
	// Create service
	srv := service.New(
		service.Name(ServerName),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterDbServiceHandler(srv.Server(), new(handler.DbService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
