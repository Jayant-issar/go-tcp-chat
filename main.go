package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listner, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listner.Close()
	log.Printf("started server on :8888")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
