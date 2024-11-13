package api

import (
	"github.com/gorilla/mux"
	"real-time-streaming/internal/api/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/stream/start", handlers.StartHandler).Methods("POST")
	router.HandleFunc("/stream/{stream_id}/send", handlers.SendHandler).Methods("POST")
	router.HandleFunc("/stream/{stream_id}/results", handlers.ResultsHandler).Methods("GET")

	return router
}
