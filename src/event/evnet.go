package event

import (
	"github.com/gorilla/mux"
	"net/http"
)

type EventServiceHandler struct{}

func (eh *EventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *EventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}
func (eh *EventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}

func ServeAPI(endpoint string) error {
	handler := &EventServiceHandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}
