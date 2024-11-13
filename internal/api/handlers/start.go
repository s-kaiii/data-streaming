package handlers

import (
	"net/http"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stream started"))
}
