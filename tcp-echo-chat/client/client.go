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

	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-client is de-activated")
	}(conn)

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
	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		conn.Write([]byte(msg))
		if msg == "quit" {
			os.Exit(0)
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
