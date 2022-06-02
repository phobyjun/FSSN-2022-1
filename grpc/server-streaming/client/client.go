package main

import (
	pb "FSSN-2022-1/grpc/server-streaming/serverstreaming"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func recvMessage(stub pb.ServerStreamingClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &pb.Number{Value: 5}
	responses, err := stub.GetServerResponse(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	for {
		response, err := responses.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[server to client] %s\n", response.Message)
	}

}

func main() {
	flag.Parse()
	channel, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer channel.Close()

	stub := pb.NewServerStreamingClient(channel)
	recvMessage(stub)
}
