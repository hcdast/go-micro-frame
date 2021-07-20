package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	dbService "github.com/hcdast/go-micro-frame/dbService/proto"
)

type DbService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *DbService) Call(ctx context.Context, req *dbService.Request, rsp *dbService.Response) error {
	log.Info("Received DbService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *DbService) Stream(ctx context.Context, req *dbService.StreamingRequest, stream dbService.DbService_StreamStream) error {
	log.Infof("Received DbService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&dbService.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *DbService) PingPong(ctx context.Context, stream dbService.DbService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&dbService.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
