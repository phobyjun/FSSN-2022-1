package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	HOST                = "127.0.0.1"
	PORT                = ":65456"
	THREAD_ACTIVE_COUNT = 0
	GROUP_QUEUE         []net.Conn
)

func main() {
	fmt.Println("> echo-server is activated")

	socket, err := net.Listen("tcp", HOST+PORT)
	CheckError(err)

	defer func(socket net.Listener) {
		recover()
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(socket)

	go mainThreadHandler()
	for {
		conn, err := socket.Accept()
		CheckError(err)
		fmt.Println("> client connected by IP address " + conn.RemoteAddr().String())
		go recvHandler(conn) // 멀티스레딩 non-blocking
	}
}

func recvHandler(conn net.Conn) {
	THREAD_ACTIVE_COUNT += 1
	GROUP_QUEUE = append(GROUP_QUEUE, conn)

	defer func(conn net.Conn) {
		THREAD_ACTIVE_COUNT -= 1
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	buffer := make([]byte, 1024)
	for {
		dlen, err := conn.Read(buffer)
		if err != nil {
			return
		}
		RecvData := string(buffer[:dlen])
		if RecvData == "quit" {
			GROUP_QUEUE = removeConnection(GROUP_QUEUE, conn)
		} else {
			fmt.Printf("> received ( %s ) and echoed to %d clients\n", RecvData, len(GROUP_QUEUE))
			for _, conn := range GROUP_QUEUE {
				_, err2 := conn.Write([]byte(RecvData))
				if err2 != nil {
					return
				}
			}
		}
	}
}

func mainThreadHandler() {
	defer func() {
		fmt.Println("> echo-server is de-activated")
		os.Exit(0)
	}()

	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		if msg == "quit" {
			if THREAD_ACTIVE_COUNT == 0 {
				fmt.Println("> stop procedure started")
				break
			} else {
				fmt.Printf("> active threads are remained: %d threads\n", THREAD_ACTIVE_COUNT)
			}
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func removeConnection(groupQueue []net.Conn, conn net.Conn) []net.Conn {
	if len(groupQueue) == 1 {
		return []net.Conn{}
	}
	for i, current := range groupQueue {
		if current == conn {
			groupQueue = append(groupQueue[:i], groupQueue[i+1:]...)
		}
	}
	return groupQueue
}
