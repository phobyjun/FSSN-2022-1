package main

import (
	"fmt"
	"log"
	"net"
)

var (
	HOST = "127.0.0.1"
	PORT = ":65456"
)

func main() {
	fmt.Println("> echo-server is activated")

	socket, err := net.Listen("tcp", HOST+PORT)
	checkError(err)

	defer func(socket net.Listener) {
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-server is de-activated")
	}(socket)

	for {
		conn, err := socket.Accept()
		checkError(err)
		fmt.Println("> client connected by IP address " + conn.RemoteAddr().String())

		buffer := make([]byte, 1024)
		dlen, err := conn.Read(buffer)
		checkError(err)

		RecvData := string(buffer[:dlen])
		fmt.Println("> echoed: " + RecvData)
		conn.Write([]byte(RecvData))
		if RecvData == "quit" {
			break
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
