package main

import (
	"io"
	"log"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	io.Copy(c, c)
}
func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "1987"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	listener, err := net.Listen("tcp", ":"+port)
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
