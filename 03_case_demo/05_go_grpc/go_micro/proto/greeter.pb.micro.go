// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/greeter.proto

package go_micro_service_imooc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for Nemo service

func NewNemoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Nemo service

type NemoService interface {
	// rpc 请求 和 响应数据格式字段定义
	SayHello(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error)
}

type nemoService struct {
	c    client.Client
	name string
}

func NewNemoService(name string, c client.Client) NemoService {
	return &nemoService{
		c:    c,
		name: name,
	}
}

func (c *nemoService) SayHello(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error) {
	req := c.c.NewRequest(c.name, "Nemo.sayHello", in)
	out := new(SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nemo service

type NemoHandler interface {
	// rpc 请求 和 响应数据格式字段定义
	SayHello(context.Context, *SayRequest, *SayResponse) error
}

func RegisterNemoHandler(s server.Server, hdlr NemoHandler, opts ...server.HandlerOption) error {
	type nemo interface {
		SayHello(ctx context.Context, in *SayRequest, out *SayResponse) error
	}
	type Nemo struct {
		nemo
	}
	h := &nemoHandler{hdlr}
	return s.Handle(s.NewHandler(&Nemo{h}, opts...))
}

type nemoHandler struct {
	NemoHandler
}

func (h *nemoHandler) SayHello(ctx context.Context, in *SayRequest, out *SayResponse) error {
	return h.NemoHandler.SayHello(ctx, in, out)
}
