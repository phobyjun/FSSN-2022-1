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

	socket, err := net.ListenPacket("udp", HOST+PORT)
	CheckError(err)

	defer func(socket net.PacketConn) {
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-server is de-activated")
	}(socket)

	for {
		recvHandler(socket)
	}
}

func recvHandler(sock net.PacketConn) {
	buffer := make([]byte, 1024)
	dlen, clientAddr, err := sock.ReadFrom(buffer)
	CheckError(err)
	RecvData := buffer[:dlen]
	fmt.Printf("> echoed: %s\n", RecvData)
	sock.WriteTo(RecvData, clientAddr)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
