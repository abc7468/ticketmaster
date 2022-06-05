package main

import (
	"sync"
	"ticketmaster/event"
)

var cnt = 0
var wg = sync.WaitGroup{}

func action() {
	cnt++
	wg.Done()
}

func main() {
	err := event.ServeAPI(":8081")
	if err != nil {
		panic(err)
	}
}
