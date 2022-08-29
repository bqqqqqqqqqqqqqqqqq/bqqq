package rpc

import (
	"google.golang.org/grpc"
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/ports"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000")
	}
	arithmeticServiceServer := grpca
	grpcSever := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcSever, arithmeticServiceServer)
	if err := grpcSever.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000:%v", err)
	}
}
