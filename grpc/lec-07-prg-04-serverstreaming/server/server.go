package main

import (
	pb "FSSN-2022-1/grpc/lec-07-prg-04-serverstreaming/serverstreaming"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedServerStreamingServer
}

func makeMessage(message string) *pb.Message {
	return &pb.Message{Message: message}
}

func (s *server) GetServerResponse(request *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	messages := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
	fmt.Printf("Server processing gRPC server-streaming {%d}.\n", request.Value)
	for _, message := range messages {
		if err := stream.Send(message); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServerStreamingServer(grpcServer, &server{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
