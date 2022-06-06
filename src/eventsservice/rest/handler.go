package rest

import (
	"net/http"
	"ticketmaster/src/lib/persistence"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func newEventHandler(databasehanlder persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehanlder,
	}
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}
