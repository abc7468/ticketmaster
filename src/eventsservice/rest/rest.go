package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"ticketmaster/src/lib/persistence"
)

func ServeAPI(endpoint, tlsendpoint string, databasehandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := newEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	go func() { httpErrChan <- http.ListenAndServe(endpoint, r) }()
	go func() { httpsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r) }()
	return httpErrChan, httpsErrChan
}
