package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, err := zmq.NewContext()
	checkError(err)
	defer ctx.Term()

	publisher, err := ctx.NewSocket(zmq.PUB)
	checkError(err)
	defer publisher.Close()
	publisher.Bind("tcp://*:5557")

	collector, err := ctx.NewSocket(zmq.PULL)
	checkError(err)
	defer collector.Close()
	collector.Bind("tcp://*:5558")

	for {
		message, err := collector.Recv(0)
		checkError(err)

		fmt.Printf("I: publishing update %s\n", message)

		publisher.Send(message, 0)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
