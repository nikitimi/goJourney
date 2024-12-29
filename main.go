package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdOn"`
}

var id = 0
var notes = make(map[string]Note)

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		panic(err)
	}

	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	notes[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetAllNoteHandler(w http.ResponseWriter, r *http.Request) {
	_notes := []Note{}
	for _, v := range notes {
		_notes = append(_notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(_notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// If `id` is provided in the request
	k := vars["id"]
	if note, ok := notes[k]; ok {
		note, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(note)
	} else {
		log.Printf("Could not find Note requested: %s", k)
	}
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	var noteToUpdate Note

	err := json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := notes[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		delete(notes, k)
		notes[k] = noteToUpdate
	} else {
		log.Printf("Could not find key of Note %s", k)
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := notes[k]; ok {
		delete(notes, k)
	} else {
		log.Printf("Could not find the key of Note %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", GetAllNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	port := ":8080"
	ListenAndServe(port, r)
}

func ListenAndServe(addr string, handler http.Handler) error {
	log.Printf("Listening to Port: %s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return server.ListenAndServe()
}
