package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	streamID := vars["stream_id"]

	// Simulate fetching results for a stream
	log.Printf("Fetching results for stream %s", streamID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Processed results for stream " + streamID))
}
