package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.10.39:4040")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("recieved response from server of", n, "bytes", "server addr: ", conn.RemoteAddr().String())
		time.Sleep(1 * time.Second)
		conn.Write([]byte("request <-->"))
	}
}
