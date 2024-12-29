package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	})
}

func main() {
	port := ":8080"
	ListenAndServe(port, messageHandler("Go Web is awesome!"))
}

func ListenAndServe(addr string, handler http.Handler) error {
	log.Printf("Listening to Port: %s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return server.ListenAndServe()
}
