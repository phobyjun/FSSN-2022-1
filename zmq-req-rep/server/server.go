package main

import (
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, err := zmq.NewContext()
	checkError(err)
	defer context.Term()

	socket, err := context.NewSocket(zmq.REP)
	checkError(err)
	defer socket.Close()

	if err := socket.Bind("tcp://*:5555"); err != nil {
		log.Fatal(err)
	}

	for {
		message, err := socket.Recv(0)
		checkError(err)

		fmt.Printf("Received request: %s\n", message)

		time.Sleep(time.Second * 1)

		socket.Send("World", 0)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
