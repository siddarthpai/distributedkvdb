package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/siddarthpai/distributedkvdb/db"
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

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value, err := db.GetKey(key)

		fmt.Fprint(w, "called get!")
		fmt.Fprintf(w, "Value = %q, error = %v", value, err)
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value := r.Form.Get("value")

		err := db.SetKey(key, []byte(value))

		fmt.Fprint(w, "called set!")
		fmt.Fprintf(w, "error = %v", err)
	})

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
