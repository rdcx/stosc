package main

import (
	"encoding/gob"
	"log"
	"net"

	"./shared"
)

func handleConnection(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	message := &shared.Message{}
	decoder.Decode(message)
	log.Printf("Received: %+v", message)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3030")
	if err != nil {
		log.Fatal("error: %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
