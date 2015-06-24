package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type Sunfish struct {
	DBSession *mgo.Session
	DB        *mgo.Database

	Router *mux.Router
	Routes []Route
}

// New returns a Sunfish object.
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

	// Create the Router
	sf.Router = newRouter(sf)
	return sf
}

func (sf *Sunfish) Close() {
	sf.DBSession.Close()
}
