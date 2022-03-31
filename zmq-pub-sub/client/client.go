package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, err := zmq.NewContext()
	checkError(err)
	defer context.Term()

	socket, err := context.NewSocket(zmq.SUB)
	checkError(err)
	defer socket.Close()

	fmt.Println("Collecting updates from weather server...")
	socket.Connect("tcp://localhost:5556")

	zipFilter := "10001"
	if len(os.Args) > 1 {
		zipFilter = os.Args[1]
	}
	socket.SetSubscribe(zipFilter)

	totalTemp := 0
	var updateNbr int
	for updateNbr = 0; updateNbr < 20; updateNbr++ {
		str, err := socket.Recv(0)
		checkError(err)
		strArray := strings.Split(str, " ")
		temperature, _ := strconv.Atoi(strArray[1])

		fmt.Printf("Receive temperature for zipcode '%s' was %d F\n", zipFilter, temperature)
	}
	fmt.Printf("Average temperature for zipcode '%s' was %d F\n", zipFilter, totalTemp/(updateNbr))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
