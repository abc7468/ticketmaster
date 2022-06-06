package main

import (
	"flag"
	"fmt"
	"log"
	"ticketmaster/src/eventsservice/rest"
	"ticketmaster/src/lib/configuration"
	"ticketmaster/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `../lib/configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
