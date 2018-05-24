package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("starting the IoT server...")
	server := NewServer()
	// run non-blocking server
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("ERROR: Failed to start server <%v>", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c  // wait for Ctrl+C
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("gracefully shutting down IoT server")
	os.Exit(0)
}
