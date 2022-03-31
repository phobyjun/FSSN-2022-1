package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending request %d...\n", i)
		s.Send("Hello", 0)

		msg, _ := s.Recv(0)
		fmt.Printf("Received reply %d [ %s ]\n", i, msg)
	}
}
