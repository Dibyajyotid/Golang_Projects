package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":2002")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(conn)

		go do(conn)
	}
}

func do(conn net.Conn) {

	buff := make([]byte, 1024)

	_, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello From The Other Side!\r\n"))
	conn.Close()
}
