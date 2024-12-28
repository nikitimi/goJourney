package main

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Go Web Development")
}

func main() {
	mux := http.NewServeMux()

	mh1 := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh1)

	http.ListenAndServe(":8080", mux)
}
