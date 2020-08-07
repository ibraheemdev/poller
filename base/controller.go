package base

import (
	"log"
	"net/http"
)

// HandleError : Writes the given status code header if the given error is not nil
func HandleError(w http.ResponseWriter, err error, code int) bool {
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return true
	}
	return false
}

// HandleNotFound :
func HandleNotFound(w http.ResponseWriter, err error) bool {
	return HandleError(w, err, http.StatusNotFound)
}

// HandleBadRequest :
func HandleBadRequest(w http.ResponseWriter, err error) bool {
	return HandleError(w, err, http.StatusBadRequest)
}
