package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	HOST = "127.0.0.1"
	PORT = ":65456"
)

func main() {
	fmt.Println("> echo-client is activated")

	conn, err := net.Dial("tcp", HOST+PORT)
	CheckError(err)

	go sendHandler(conn)
	for {
		buffer := make([]byte, 1024)
		dlen, err := conn.Read(buffer)
		CheckError(err)
		RecvData := string(buffer[:dlen])
		fmt.Println("> received: " + RecvData)
	}
}

func sendHandler(conn net.Conn) {
	defer func() {
		fmt.Println("> echo-client is de-activated")
		os.Exit(0)
	}()

	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		conn.Write([]byte(msg))
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
