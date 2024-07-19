package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
)

var (
	dbLocation = flag.String("db-location", "", "The path which leads to the bolt database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP address") //127.0.0.1:8080 is the default value
)

func parseFlags() { //function to validate
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("Provide a valid DB location")
	}
}
func main() {
	parseFlags()
	db, err := bolt.Open(*dbLocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "called get!")
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "called set!")
	})

	log.Fatal(http.ListenAndServe(*httpAddr, nil))

}
