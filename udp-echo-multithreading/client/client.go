package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	HOST        = "127.0.0.1"
	PORT        = ":"
	SERVER_PORT = ":65456"
)

func main() {
	fmt.Println("> echo-client is activated")

	socket, err := net.ListenPacket("udp", HOST+PORT)
	CheckError(err)

	server, err := net.ResolveUDPAddr("udp", HOST+SERVER_PORT)
	CheckError(err)

	defer func(socket net.PacketConn) {
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-client is de-activated")
	}(socket)

	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		socket.WriteTo([]byte(msg), server)

		buffer := make([]byte, 1024)
		dlen, _, err := socket.ReadFrom(buffer)
		CheckError(err)
		RecvData := buffer[:dlen]
		fmt.Printf("> received: %s\n", RecvData)

		if msg == "quit" {
			break
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
