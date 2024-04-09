package service

import (
	"context"

	"github.com/fanchann/grpc_math/proto"
)

type mathServiceImpl struct {
	proto.MathServiceServer
}

func NewMathServiceImpl() proto.MathServiceServer {
	return &mathServiceImpl{}
}

func (m *mathServiceImpl) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	result := req.GetA() + req.GetB()

	return &proto.Response{Result: result}, nil
}

func (m *mathServiceImpl) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	result := req.GetA() * req.GetB()

	return &proto.Response{Result: result}, nil
}

func (m *mathServiceImpl) Reduce(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	result := req.GetA() - req.GetB()

	return &proto.Response{Result: result}, nil
}

func (m *mathServiceImpl) Devide(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	result := req.GetA() / req.GetB()

	return &proto.Response{Result: result}, nil
}
