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

	socket, err := context.NewSocket(zmq.REQ)
	checkError(err)
	defer socket.Close()

	if err := socket.Connect("tcp://localhost:5555"); err != nil {
		log.Fatal(err)
	}

	for request := 0; request < 10; request++ {
		fmt.Printf("Seding request %d ...\n", request)
		socket.Send("Hello", 0)

		message, err := socket.Recv(0)
		checkError(err)

		fmt.Printf("Received reply %d [ %s ]\n", request, message)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
