// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: videoService.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for VideoService service

func NewVideoServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VideoService service

type VideoService interface {
	UploadVideo(ctx context.Context, in *VideoRequest, opts ...client.CallOption) (*VideoResponse, error)
}

type videoService struct {
	c    client.Client
	name string
}

func NewVideoService(name string, c client.Client) VideoService {
	return &videoService{
		c:    c,
		name: name,
	}
}

func (c *videoService) UploadVideo(ctx context.Context, in *VideoRequest, opts ...client.CallOption) (*VideoResponse, error) {
	req := c.c.NewRequest(c.name, "VideoService.UploadVideo", in)
	out := new(VideoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VideoService service

type VideoServiceHandler interface {
	UploadVideo(context.Context, *VideoRequest, *VideoResponse) error
}

func RegisterVideoServiceHandler(s server.Server, hdlr VideoServiceHandler, opts ...server.HandlerOption) error {
	type videoService interface {
		UploadVideo(ctx context.Context, in *VideoRequest, out *VideoResponse) error
	}
	type VideoService struct {
		videoService
	}
	h := &videoServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VideoService{h}, opts...))
}

type videoServiceHandler struct {
	VideoServiceHandler
}

func (h *videoServiceHandler) UploadVideo(ctx context.Context, in *VideoRequest, out *VideoResponse) error {
	return h.VideoServiceHandler.UploadVideo(ctx, in, out)
}
