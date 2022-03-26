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
		err := socket.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("> echo-server is de-activated")
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
			removeConnection(&GROUP_QUEUE, conn)
		}
		fmt.Printf("> received (%s) and echoed to %d clients", RecvData, len(GROUP_QUEUE))
		for _, conn := range GROUP_QUEUE {
			_, err2 := conn.Write([]byte(RecvData))
			if err2 != nil {
				return
			}
		}
	}
}

func mainThreadHandler() {
	for {
		fmt.Print("> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg := sc.Text()
		if msg == "quit" {
			if THREAD_ACTIVE_COUNT == 0 {
				fmt.Println("> stop procedure started")
				os.Exit(0)
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

func removeConnection(groupQueue *[]net.Conn, conn net.Conn) {
	queue := *groupQueue
	for i, current := range queue {
		if current == conn {
			queue = append(queue[:i], queue[i+1:]...)
		}
	}
	groupQueue = &queue
}
