package rpc

import (
	"context"
	"hex/internal/adapters/framework/left/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpcAdaptor Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcAdaptor.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcAdaptor Adapter) GetSubstraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcAdaptor.api.GetSubstraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcAdaptor Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcAdaptor.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcAdaptor Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcAdaptor.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
