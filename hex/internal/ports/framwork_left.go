package ports

import (
	"context"
	"hex/internal/adapters/framework/left/gRPC/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, rq *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, rq *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, rq *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, rq *pb.OperationParameters) (*pb.Answer, error)
}