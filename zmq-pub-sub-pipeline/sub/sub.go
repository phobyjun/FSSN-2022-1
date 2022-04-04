package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, err := zmq.NewContext()
	checkError(err)
	defer ctx.Term()

	subscriber, err := ctx.NewSocket(zmq.SUB)
	checkError(err)
	defer subscriber.Close()
	subscriber.SetSubscribe("")
	subscriber.Connect("tcp://localhost:5557")

	publisher, err := ctx.NewSocket(zmq.PUSH)
	checkError(err)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5558")

	poller := zmq.NewPoller()
	poller.Add(subscriber, zmq.POLLIN)

	rand.Seed(time.Now().UnixNano())
	for {
		if p, _ := poller.Poll(100 * time.Millisecond); len(p) > 0 {
			message, err := subscriber.Recv(0)
			checkError(err)
			fmt.Printf("I: received message %s\n", message)
		} else {
			randNum := rand.Intn(100-1) + 1
			if randNum < 10 {
				publisher.Send(strconv.Itoa(randNum), 0)
				fmt.Printf("I: sending message %d\n", randNum)
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
