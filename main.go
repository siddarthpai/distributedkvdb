package main

import (
	"flag"
	"log"

	bolt "go.etcd.io/bbolt"
)

var (
	dbLocation = flag.String("db-location", "", "The path which leads to the bolt database")
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
}
