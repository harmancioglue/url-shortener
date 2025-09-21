package main

import (
	"context"
	"harmancioglue/url-shortener/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("🚀 Url Shortener Service starting...")

	config, err := config.Load()
	if err != nil {
		log.Fatalf("Config creation error: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

}
