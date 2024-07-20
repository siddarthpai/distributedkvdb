package db

import (
	bolt "go.etcd.io/bbolt"
)

type Database struct {
	db *bolt.DB
}

var defaultBucket = []byte("default value")

// This is a constructor which will return an instance of a database we can work with
func NewDatabase(dbPath string) (db *Database, closeFunc func() error, err error) { // the closeFunc has a return type of error and that is why we assign error type to it
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, nil, err // this is a better approach than using Fatal(provided in the help doc)
	}

	closeFunc = boltDb.Close

	return &Database{db: boltDb}, closeFunc, nil
}

// sets the key to the requested value(which is in the form of bytes) into the default database else will return an error
func (d *Database) SetKey(key string, value []byte) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		bu := tx.Bucket(defaultBucket)
		return bu.Put([]byte(key), value)
	})
}

// gets the value for the requested key else will return an error
func (d *Database) GetKey(key string) ([]byte, error) {
	var result []byte

	err := d.db.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket(defaultBucket)
		result = bu.Get([]byte(key))
		return nil
	})

	if err == nil {
		return result, nil
	}

	return nil, err
}
