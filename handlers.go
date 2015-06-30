package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Auth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO Handle auth.
		// Only continue if user is authorized
		fn(w, r)
	}
}

// AddFile handles a Post Request for a Sia file and saves it to the DB
func (sf *Sunfish) AddFile(w http.ResponseWriter, r *http.Request) {
	var siafile Siafile
	const maxSiaFilesize = 1 << 20 // 1 MiB
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxSiaFilesize))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &siafile); err != nil {
		// If cannot process the Siafile
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	// Handle saving the Siafile to the db and file system
	siafile.UploadedTime = time.Now()
	err = sf.DB.C("siafiles").Insert(siafile)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		panic(err)
	}
}

//  GetAll returns all siafiles as a json list.
func (sf *Sunfish) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO Pagination and sorts
	var siafiles []Siafile

	// Select removes the content from query results use for not returning .sia
	err := sf.DB.C("siafiles").Find(bson.M{}).All(&siafiles)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafiles); err != nil {
		panic(err)
	}
}

// GetFile takes the hash of a Siafile in the URL and returns the Siafile in JSON
func (sf *Sunfish) GetFile(w http.ResponseWriter, r *http.Request) {
	var id string
	var siafile Siafile

	vars := mux.Vars(r)
	id = vars["id"]

	// Query and find by one id
	err := sf.DB.C("siafiles").FindId(bson.ObjectIdHex(id)).One(&siafile)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		panic(err)
	}
}

// SearchFile takes a query parameter of tags and searches the DB for siafiles
// that match that tag.
func (sf *Sunfish) SearchFile(w http.ResponseWriter, r *http.Request) {
	var siafiles []Siafile
	var search string

	query := r.URL.Query()
	search = query.Get("tags")
	// Searches db or all siafiles that have the query string in it's tags
	err := sf.DB.C("siafiles").Find(bson.M{"tags": search}).All(&siafiles)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafiles); err != nil {
		panic(err)
	}
}

// Takes the hash of a Siafile in the URL and removes it from DB
func (sf *Sunfish) DeleteFile(w http.ResponseWriter, r *http.Request) {
	var id string
	//var siafile Siafile

	vars := mux.Vars(r)
	id = vars["id"]

	// Db.C("siafiles").remove({'hash': hash})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
}
