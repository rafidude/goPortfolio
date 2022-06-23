package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// NewDatabase -
func NewDatabase(path string) *Database {
	// If the path already exists, skip the error, that's okay
	// If the path doesn't exist, i.e no error, create it.
	if err := os.Mkdir(path, 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}
	return &Database{path}
}

// Database -
type Database struct {
	path string
}

func (d *Database) getKey(key string) string {
	return fmt.Sprintf("%s/%s", d.path, key)
}

// Write -
func (d *Database) Write(key string, value interface{}) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(d.getKey(key), encoded, 0666); err != nil {
		return err
	}

	return nil
}

// Get -
func (d *Database) Get(key string) (interface{}, error) {
	data, err := ioutil.ReadFile(d.getKey(key))
	if err != nil {
		return nil, err
	}

	var r interface{}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	db := NewDatabase(".data")

	if err := db.Write("key", "value"); err != nil {
		log.Panic(err)
	}

	val, err := db.Get("key")
	if err != nil {
		log.Panic(err)
	}

	log.Println(val)
}
