package main

import (
	"log"
	"net"
)

func main() {
	listener, err := createListener("127.0.0.1", 3000)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	log.Printf("listening on %q", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			defer func() { _ = c.Close() }()
			log.Printf("new connection from %q", c.RemoteAddr())
		}(conn)
	}
}
