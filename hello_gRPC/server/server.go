package main

import (
	"FSSN-2022-1/hello_gRPC"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "FSSN-2022-1/hello_gRPC/hellogrpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyFunction(ctx context.Context, request *pb.MyNumber) (*pb.MyNumber, error) {
	response := pb.MyNumber{
		Value: hello_gRPC.MyFunc(request.Value),
	}
	return &response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
