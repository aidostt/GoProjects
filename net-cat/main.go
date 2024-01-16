package main

import (
	"fmt"
	"log"
	"net-cat/server"
	"os"
)

const (
	usage       = "[USAGE]: ./TCPChat $port"
	defaultPort = "8989"
)

func main() {
	port := defaultPort
	switch len(os.Args) {
	case 2:
		port = os.Args[1]
	case 1:
		// default port
	default:
		log.Println(usage)
		return
	}

	s := server.New(fmt.Sprintf(":%v", port), 10)
	s.Start()
}
