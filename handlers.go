package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Hello from Sevalla")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
		log.Printf("failed to write health response: %v", err)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprintf(w, `{"version":"%s","commit":"%s","built":"%s"}`+"\n", buildVersion, buildCommit, buildTime); err != nil {
		log.Printf("failed to write version response: %v", err)
	}
}
