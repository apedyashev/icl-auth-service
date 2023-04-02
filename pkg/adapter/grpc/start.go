package grpc

import (
	"fmt"
	"icl-auth/pkg/adapter/grpc/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartAllGrpcServers(db *gorm.DB) {
	gRpcPort := os.Getenv("GRPC_PORT")
	if gRpcPort == "" {
		gRpcPort = "50001"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, NewUserServiceServer(db))
	log.Printf("gRPC is started on port %s\n", gRpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to serve: %v\n", err)
	}
}
