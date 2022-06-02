package main

import (
	pb "FSSN-2022-1/grpc/bidirectional-streaming/bidirectional"
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

func makeMessage(message string) *pb.Message {
	return &pb.Message{Message: message}
}

func generateMessage() []*pb.Message {
	messages := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}

	return messages
}

func sendMessage(stub pb.BidirectionalClient) {
	messages := generateMessage()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	responses, err := stub.GetServerResponse(ctx)
	if err != nil {
		log.Fatal(err)
	}

	wait := make(chan struct{})
	go func() {
		for {
			in, err := responses.Recv()
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[server to client] %s\n", in.Message)
		}
	}()

	for _, message := range messages {
		if err := responses.Send(message); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[client to server] %s\n", message.Message)
	}
	responses.CloseSend()
	<-wait
}

func main() {
	flag.Parse()
	channel, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer channel.Close()

	stub := pb.NewBidirectionalClient(channel)
	sendMessage(stub)
}
