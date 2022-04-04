package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, err := zmq.NewContext()
	checkError(err)
	defer context.Term()

	publisher, err := context.NewSocket(zmq.PUB)
	checkError(err)
	defer publisher.Close()

	collector, err := context.NewSocket(zmq.PULL)
	checkError(err)
	defer collector.Close()

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
