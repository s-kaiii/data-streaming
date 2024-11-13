package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SendHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	streamID := vars["stream_id"]

	// Simulate receiving data chunks
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	data := r.FormValue("data")

	log.Printf("Received data for stream %s: %s", streamID, data)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Data received"))
}
