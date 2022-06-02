package main

import (
	pb "FSSN-2022-1/grpc/lec-07-prg-03-clientstreaming/clientstreaming"
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
	pb.UnimplementedClientStreamingServer
}

func (s *server) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	var count int32 = 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Number{
				Value: count,
			})
		}
		if err != nil {
			return err
		}
		count += 1
	}

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterClientStreamingServer(grpcServer, &server{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
