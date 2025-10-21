package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	serverPort        = ":8000"
	shutdownTimeout   = 10 * time.Second
	readHeaderTimeout = 2 * time.Second
	readTimeout       = 5 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 60 * time.Second
	maxHeaderBytes    = 1 << 20 // 1MB
)

var newRunContext = func() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
}

func main() {
	ctx, cancel := newRunContext()
	defer cancel()

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/version", version)

	// Apply middleware
	handler := chain(mux, recoverMiddleware, loggingMiddleware)

	// Configure server
	srv := &http.Server{
		Addr:              serverPort,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	// Start server in background
	go func() {
		log.Printf("Listening on port %s", serverPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	log.Println("shutdown signal received")

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}
