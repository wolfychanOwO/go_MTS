package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	srv "main/task2/server"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/version", srv.VersionHandler)
	http.HandleFunc("/decode", srv.DecodeHandler)
	http.HandleFunc("/hard-op", srv.HardOpHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
