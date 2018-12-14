package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "grpcHeratbeat/services"
)

const (
	address = "localhost:25000"
)

// a simple grpc client to connect to the GRPC server for heartbeat
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc endpoint: %v", err)
	}
	defer conn.Close()

	// create the client instance
	client := pb.NewHeartbeatClient(conn)

	// connect to the server and call the required method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetHeartbeat(ctx, &pb.Empty{})
	//c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Error in calling the heartbeat service : %v", err)
	}
	log.Printf("Service heartbeat received : %s", response.Message)
}
