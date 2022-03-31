package main

import (
	"fmt"
	"log"
	"math/rand"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("Publishing updates at weather server...")

	context, err := zmq.NewContext()
	checkError(err)
	defer context.Term()

	socket, err := context.NewSocket(zmq.PUB)
	checkError(err)
	defer socket.Close()

	socket.Bind("tcp://*:5556")

	for {
		zipcode := rand.Intn(100000-1) + 1
		temperature := rand.Intn(135+80) - 80
		relhumidity := rand.Intn(60-10) + 10

		socket.Send(fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity), 0)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
