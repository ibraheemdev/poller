package polls

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Create : POST "/polls"
func Create() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		poll := new(PollParams)
		err := json.NewDecoder(r.Body).Decode(&poll)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pollID, errs := createPoll(poll)
		strErrors := make([]string, len(errs))
		for i, err := range errs {
			strErrors[i] = err.Error()
		}
		if errs != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(strErrors)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"_id":%s}`, pollID)))
	}
}

// Show : GET "/polls/:id"
func Show() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		poll := &Poll{}
		err := poll.Collection().Find(ps.ByName("id"), poll)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json, err := json.Marshal(poll)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(json)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Update : PUT "/polls/:id"
func Update() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pollID := ps.ByName("id")
		poll := new(PollParams)
		err := json.NewDecoder(r.Body).Decode(&poll)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updatePoll(pollID, poll)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
