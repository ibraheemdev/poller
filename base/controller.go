package base

import (
	"log"
	"net/http"
)

// HandleNotFound : Returns whether the error was not nil
func HandleNotFound(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return true
	}
	return false
}

// HandleBadRequest : Returns whether the error was not nil
func HandleBadRequest(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	return false
}
