package main

import (
	"context"
	"fmt"
	"harmancioglue/url-shortener/internal/api/http"
	"harmancioglue/url-shortener/internal/app"
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

	application, err := app.Init(config)
	if err != nil {
		log.Fatalf("Application creation error: %v", err)
	}

	api := http.NewApi(application)

	go func() {
		addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		if err = api.Server.Listen(addr); err != nil {
			log.Println("🔥 Url Shortener Service closing...")
			ch <- os.Interrupt
		}
	}()

	go func() {
		<-ch
		cancel()
	}()

	<-ctx.Done()

	err = api.Server.Shutdown()

}
