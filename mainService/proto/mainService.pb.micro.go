// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mainService.proto

package mainService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MainService service

func NewMainServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MainService service

type MainService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (MainService_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (MainService_PingPongService, error)
}

type mainService struct {
	c    client.Client
	name string
}

func NewMainService(name string, c client.Client) MainService {
	return &mainService{
		c:    c,
		name: name,
	}
}

func (c *mainService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "MainService.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (MainService_StreamService, error) {
	req := c.c.NewRequest(c.name, "MainService.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &mainServiceStream{stream}, nil
}

type MainService_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type mainServiceStream struct {
	stream client.Stream
}

func (x *mainServiceStream) Close() error {
	return x.stream.Close()
}

func (x *mainServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *mainServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mainServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mainServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mainService) PingPong(ctx context.Context, opts ...client.CallOption) (MainService_PingPongService, error) {
	req := c.c.NewRequest(c.name, "MainService.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &mainServicePingPong{stream}, nil
}

type MainService_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type mainServicePingPong struct {
	stream client.Stream
}

func (x *mainServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *mainServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *mainServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mainServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mainServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *mainServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MainService service

type MainServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, MainService_StreamStream) error
	PingPong(context.Context, MainService_PingPongStream) error
}

func RegisterMainServiceHandler(s server.Server, hdlr MainServiceHandler, opts ...server.HandlerOption) error {
	type mainService interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type MainService struct {
		mainService
	}
	h := &mainServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MainService{h}, opts...))
}

type mainServiceHandler struct {
	MainServiceHandler
}

func (h *mainServiceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.MainServiceHandler.Call(ctx, in, out)
}

func (h *mainServiceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MainServiceHandler.Stream(ctx, m, &mainServiceStreamStream{stream})
}

type MainService_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type mainServiceStreamStream struct {
	stream server.Stream
}

func (x *mainServiceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *mainServiceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *mainServiceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mainServiceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mainServiceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *mainServiceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.MainServiceHandler.PingPong(ctx, &mainServicePingPongStream{stream})
}

type MainService_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type mainServicePingPongStream struct {
	stream server.Stream
}

func (x *mainServicePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *mainServicePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *mainServicePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mainServicePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mainServicePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *mainServicePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
