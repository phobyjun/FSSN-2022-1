package main

import (
	pb "FSSN-2022-1/hello_gRPC/hellogrpc"
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	channel, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer channel.Close()
	stub := pb.NewMyServiceClient(channel)

	request := pb.MyNumber{
		Value: 4,
	}
	response, err := stub.MyFunction(context.Background(), &request)
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	log.Printf("gRPC result: %d", response.Value)
}
