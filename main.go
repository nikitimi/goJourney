package main

import (
	"fmt"
	"net/http"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	})
}

func main() {
	mux := http.NewServeMux()

	mh1 := messageHandler("Go Web is awesome!")
	mux.Handle("/welcome", mh1)

	http.ListenAndServe(":8080", mux)
}
