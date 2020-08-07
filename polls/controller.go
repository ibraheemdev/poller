package polls

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ibraheemdev/poller/base"
	"github.com/julienschmidt/httprouter"
)

// Create : POST "/polls"
func Create() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		poll := new(PollParams)
		err := json.NewDecoder(r.Body).Decode(&poll)
		if badReq := base.HandleBadRequest(w, err); badReq {
			return
		}
		pollID, errs := createPoll(poll)
		if errs != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(errs)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"_id":%s}`, pollID)))
	}
}

// Show : GET "/polls/:id"
func Show() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		poll := new(Poll)
		err := poll.Collection().Find(ps.ByName("id"), poll)
		if nf := base.HandleNotFound(w, err); nf {
			return
		}
		json, err := json.Marshal(poll)
		if badReq := base.HandleBadRequest(w, err); badReq {
			return
		}
		if _, err = w.Write(json); err != nil {
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
		if badReq := base.HandleBadRequest(w, err); badReq {
			return
		}
		err = updatePoll(pollID, poll)
		if badReq := base.HandleBadRequest(w, err); badReq {
			return
		}
	}
}
