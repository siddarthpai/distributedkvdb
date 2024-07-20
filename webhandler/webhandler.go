package webhandler

import (
	"fmt"
	"net/http"

	"github.com/siddarthpai/distributedkvdb/db"
)

type Server struct {
	db *db.Database
}

func NewServer(db *db.Database) *Server { //Constructor to create a new server instance with HTTP handlers for getting and setting
	return &Server{
		db: db,
	}
}

// handling the get endpoint
func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value, err := s.db.GetKey(key)

	fmt.Fprint(w, "called get!")
	fmt.Fprintf(w, "Value = %q, error = %v", value, err)
}

// handling the set endpoint
func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")
	err := s.db.SetKey(key, []byte(value))

	fmt.Fprint(w, "called set!")
	fmt.Fprintf(w, "error = %v", err)
}

// function to listen and serve for the server
func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
