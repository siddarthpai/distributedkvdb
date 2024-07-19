package db

import (
	bolt "go.etcd.io/bbolt"
)

type Database struct {
	db *bolt.DB
}

// This is a constructor which will return an instance of a database we can work with
func NewDatabase(dbPath string) (db *Database, closeFunc func() error, err error) { // the closeFunc has a return type of error and that is why we assign error type to it
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, nil, err // this is a better approach than using Fatal(provided in the help doc)
	}

	closeFunc = boltDb.Close

	return &Database{db: boltDb}, closeFunc, nil
}
