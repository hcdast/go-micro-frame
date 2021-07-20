package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	mainService "github.com/hcdast/go-micro-frame/mainService/proto"
)

type MainService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MainService) Call(ctx context.Context, req *mainService.Request, rsp *mainService.Response) error {
	log.Info("Received MainService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MainService) Stream(ctx context.Context, req *mainService.StreamingRequest, stream mainService.MainService_StreamStream) error {
	log.Infof("Received MainService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&mainService.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MainService) PingPong(ctx context.Context, stream mainService.MainService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&mainService.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
