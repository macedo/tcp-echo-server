package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	io.Copy(c, c)
}
func main() {
	listener, err := net.Listen("tcp", "localhost:1987")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
