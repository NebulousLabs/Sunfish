package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"path/filepath"
)

// Sunfish object is the main server object storing routes and the connection
// to the database.
type Sunfish struct {
	DBSession *mgo.Session
	DB        *mgo.Database

	Router *mux.Router
	Routes []Route

	logger  *log.Logger
	logFile *os.File
}

func NewLogger(sf *Sunfish, logDir string) {
	// Make the log directory
	err := os.MkdirAll(logDir, 0700)
	if err != nil {
		os.Exit(1)
	}

	logFile, err := os.OpenFile(filepath.Join(logDir, "sunfish.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		os.Exit(1)
	}
	sf.logFile = logFile
	sf.logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	return
}

// NewSunfish returns a Sunfish object.
func NewSunfish(logDir string, dbName string) *Sunfish {
	sf := new(Sunfish)

	// Create the Database
	var err error
	sf.DBSession, err = mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Could not reach a Mongo server. Make you have Mongo configured correctly.")
		os.Exit(1)
	}
	sf.DBSession.SetMode(mgo.Monotonic, true)
	sf.DB = sf.DBSession.DB(dbName)

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

	NewLogger(sf, logDir)

	sf.logger.Println("INFO: Succesfully created a new Sunfish Object.")
	return sf
}

// Close cleans up the sunfish object. Needed to close db connection and any other
// shutdown tasks
func (sf *Sunfish) Close() {
	sf.DBSession.Close()
}
