package main

import (
	"net/http"
	"log"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Response with 5xx error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	//implement respondWithJSON
	responseWithJSON(w, code, errorResponse{Error: msg})
}
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"message": "Ready",
	}
	responseWithJSON(w, http.StatusOK, payload)
}
