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

	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		conn.Write([]byte(msg))
		CheckError(err)

		buffer := make([]byte, 1024)
		dlen, err := conn.Read(buffer)
		CheckError(err)
		RecvData := string(buffer[:dlen])
		fmt.Println("> received: " + RecvData)
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
