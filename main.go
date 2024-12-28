package main

import (
	"fmt"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandler{"This is awesome!"}
	mux.Handle("/message", mh2)
	// fs := http.FileServer(http.Dir("public"))
	http.ListenAndServe(":8080", mux)
}
