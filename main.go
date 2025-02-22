package main

import (
	"log"
)

func main() {
	listener, err := createListener("127.0.0.1", 3000)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	log.Printf("listening on %q", listener.Addr())
}
