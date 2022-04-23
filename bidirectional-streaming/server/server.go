package main

import (
	pb "FSSN-2022-1/bidirectional-streaming/bidirectional"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedBidirectionalServer
}

func (s *server) GetServerResponse(stream pb.Bidirectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC bidirectional streaming.")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		message := in.GetMessage()
		stream.SendMsg(message)
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBidirectionalServer(grpcServer, &server{})
	fmt.Println("Starting server. Listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
