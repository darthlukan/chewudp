package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	srv, err := net.ResolveUDPAddr("udp", ":65432")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", srv)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		size, from, err := conn.ReadFromUDP(buffer)

		if err != nil {
			panic(err)
		}

		if size > 0 {
			fmt.Printf("Received: %v From: %v", string(buffer[:size]), from)
		}
		fmt.Println("Waiting for packets...")
		os.Stdout.Sync()
		time.Sleep(1000 * time.Millisecond)
	}
}
