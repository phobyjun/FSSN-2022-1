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
	CheckError(err)

	defer func(socket net.Listener) {
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-server is de-activated")
	}(socket)

	con, err := socket.Accept()
	CheckError(err)
	fmt.Println("> client connected by IP address " + con.RemoteAddr().String())

	for {
		buffer := make([]byte, 1024)
		dlen, err := con.Read(buffer)
		CheckError(err)

		RecvData := string(buffer[:dlen])
		fmt.Println("> echoed: " + RecvData)
		con.Write([]byte(RecvData))
		if RecvData == "quit" {
			break
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
