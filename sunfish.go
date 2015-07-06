package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
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
		panic(err)
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

	// Ensure index make sure that selected fields are indexed by mongo
	err = sf.DB.C("siafiles").EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Create the Router
	sf.Router = newRouter(sf)
	return sf
}

// Close cleans up the sunfish object. Needed to close db connection and any other
// shutdown tasks
func (sf *Sunfish) Close() {
	sf.DBSession.Close()
}
