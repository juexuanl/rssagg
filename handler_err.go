package main

import (	
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	// Example usage of responseWithError
	responseWithError(w, http.StatusBadRequest, "An error occurred")
}

