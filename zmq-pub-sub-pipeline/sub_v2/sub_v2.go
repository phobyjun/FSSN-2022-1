package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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

	clientID := os.Args[1]
	rand.Seed(time.Now().UnixNano())
	for {
		if p, _ := poller.Poll(100 * time.Millisecond); len(p) > 0 {
			message, err := subscriber.Recv(0)
			checkError(err)
			fmt.Printf("%s: receive status => %s\n", clientID, message)
		} else {
			randNum := rand.Intn(100-1) + 1
			if randNum < 10 {
				time.Sleep(1 * time.Second)
				msg := fmt.Sprintf("(%s:ON)", clientID)
				publisher.Send(msg, 0)
				fmt.Printf("%s: send status - activated\n", clientID)
			} else if randNum > 90 {
				time.Sleep(1 * time.Second)
				msg := fmt.Sprintf("(%s:OFF)", clientID)
				publisher.Send(msg, 0)
				fmt.Printf("%s: send status - deactivated\n", clientID)
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
