package main

import (
	"FSSN-2022-1/grpc/lec-07-prg-01-hello_gRPC"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "FSSN-2022-1/grpc/lec-07-prg-01-hello_gRPC/hellogrpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyFunction(ctx context.Context, request *pb.MyNumber) (*pb.MyNumber, error) {
	response := pb.MyNumber{
		Value: lec_07_prg_01_hello_gRPC.MyFunc(request.Value),
	}
	return &response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, &server{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
