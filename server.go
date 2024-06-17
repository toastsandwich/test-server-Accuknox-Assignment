package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "192.168.10.39:4040")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("connection recieved\n"))
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("recieved request of", n, "bytes client addr: ", conn.RemoteAddr().String())
		conn.Write([]byte("response to request"))
		time.Sleep(1 * time.Second)
	}
}
