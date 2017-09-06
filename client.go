package main

import (
	"encoding/gob"
	"log"
	"net"

	"./shared"
)

func main() {
	log.Println("Creating client")
	conn, err := net.Dial("tcp", "localhost:3030")
	if err != nil {
		log.Fatal("error: %v", err)
	}
	encode := gob.NewEncoder(conn)
	msg := &shared.Message{"Hello it's me your client"}
	encode.Encode(msg)
	conn.Close()
	log.Println("Message sent")
}
