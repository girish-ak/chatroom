package main

import (
	"log"
	"net"
)

func main() {

	s := newServer()
	go s.run()
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("Server can't be started : %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Server started on -> 7070")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection can't be accepted : %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
