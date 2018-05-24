package main

import (
	"log"
	"os"
)

func main() {
	log.Println("starting the IoT server...")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3027"
	}
	server := NewServer()
	server.Run(":" + port)
}
