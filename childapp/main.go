package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/nrnrk/bazel-sample/hello/hello"
)

const (
	port = ":55550"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &HelloServer{})
	log.Printf("listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type HelloServer struct {
	pb.UnimplementedHelloServer
}

func (s *HelloServer) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
