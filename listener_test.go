package main

import (
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := createListener("127.0.0.1", 3000)
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	t.Logf("listening on %s", listener.Addr())
}
