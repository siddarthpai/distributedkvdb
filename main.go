package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/siddarthpai/distributedkvdb/db"
	"github.com/siddarthpai/distributedkvdb/webhandler"
)

var (
	dbLocation = flag.String("db-location", "", "The path which leads to the bolt database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP address") // 127.0.0.1:8080 is the default value
)

func parseFlags() { // function to validate
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("Provide a valid DB location")
	}
}

func main() {
	parseFlags()

	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("error creating NewDatabase(%q): %v", *dbLocation, err)
	}
	defer close()

	server := webhandler.NewServer(db)
	http.HandleFunc("/get", server.GetHandler)

	http.HandleFunc("/set", server.SetHandler)

	log.Fatal(server.ListenAndServe(*httpAddr))
}
