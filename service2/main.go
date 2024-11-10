package main

import (
    "context"
    "log"
    "net"

    pb "github.com/hamidehalaee/proto/github.com/hamidehalaee/proto/example" // Import the shared proto package
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterExampleServiceServer(grpcServer, &server{})
    log.Println("service2 is running on port :50052")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
