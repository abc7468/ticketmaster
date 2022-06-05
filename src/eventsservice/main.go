package main

import (
	"ticketmaster/eventsservice/rest"
)

func main() {
	err := rest.ServeAPI(":8081")
	if err != nil {
		panic(err)
	}
}
