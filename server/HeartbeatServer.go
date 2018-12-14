package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "grpcHeratbeat/services"
)

//go:generate protoc -I ../services/ ../services/heartbeat.proto --go_out=plugins=grpc:../services

const (
	port = ":25000"
)

// server is used to implement api.HeartbeatServer
type server struct{}

// SayHello implements api.HeartbeatServer
func (s *server) GetHeartbeat(ctx context.Context, in *pb.Empty) (*pb.HeartbeatMsg, error) {
	//TODO add any service specific code determine the tatus and return the results
	return &pb.HeartbeatMsg{ServiceName: "sample1", Timestamp: 1544823909, IsRunning: true, Message: "500 OK"}, nil
}

// Runs the gRPC server
func main() {
	// open the port and start listening to it
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port : %v", err)
	}
	// create a new grpc server instance
	grpcServer := grpc.NewServer()

	//register the heartbeat server
	pb.RegisterHeartbeatServer(grpcServer, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// bing the port through which clients cn connect
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to expose the end point: %v", err)
	}
}
