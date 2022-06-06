package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"ticketmaster/src/lib/persistence"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := newEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}
