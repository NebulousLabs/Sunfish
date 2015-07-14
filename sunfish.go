package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"os"
)

// Sunfish object is the main server object storing routes and the connection
// to the database.
type Sunfish struct {
	DBSession *mgo.Session
	DB        *mgo.Database

	Router *mux.Router
	Routes []Route
}

// NewSunfish returns a Sunfish object.
func NewSunfish() *Sunfish {
	sf := new(Sunfish)

	// Create the Database
	var err error
	sf.DBSession, err = mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Could not reach a Mongo server. Ensure Mongo is configured correctly.")
		os.Exit(1)
	}
	sf.DBSession.SetMode(mgo.Monotonic, true)
	sf.DB = sf.DBSession.DB("sunfish")

	// Index tags and title fields for faster searching
	index := mgo.Index{
		Key:        []string{"tags", "title"},
		Unique:     false,
		DropDups:   false,
		Background: false,
		Sparse:     true,
	}

	// Ensure selected fields are indexed by mongo
	err = sf.DB.C("siafiles").EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Create the Router
	sf.Router = newRouter(sf)
	return sf
}

// Close cleans up the sunfish object's db connection and other shutdown tasks
func (sf *Sunfish) Close() {
	sf.DBSession.Close()
}
