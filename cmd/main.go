package main

import (
	"fmt"
	"github.com/mmarci96/reverse-proxy-go/internal/server"
	"log"
)

func main() {
	fmt.Println("Reverse proxy starting...")
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
	select {}
}

