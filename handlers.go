package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
		sf.logger.Println("ERROR: Could not read body of request.")
		return
	}

	if err := r.Body.Close(); err != nil {
		sf.logger.Println("ERROR: Could not close body of request.")
		return
	}

	if err := json.Unmarshal(body, &siafile); err != nil {
		// If cannot process the Siafile
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			sf.logger.Println("ERROR: Could not encode json response.")
			return
		}
		sf.logger.Println("INFO: Could not process the uploaded Siafile.")
		return
	}

	// Validate fields
	var errors []string
	// Filename has to contain '.sia'
	if !strings.Contains(siafile.Filename, ".sia") || len(siafile.Filename) == 0 {
		errors = append(errors, "Bad Siafile upload")
	}
	if len(siafile.Title) == 0 {
		errors = append(errors, "Title field can not be blank")
	}
	for i, tag := range siafile.Tags {
		siafile.Tags[i] = strings.ToLower(tag)
	}

	if len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(errors); err != nil {
			sf.logger.Println("ERROR: Siafile was not valid.")
			return
		}
		return
	}

	// Handle saving the Siafile to the db and file system
	siafile.UploadedTime = time.Now()
	siafile.Id = bson.NewObjectId()
	err = sf.DB.C("siafiles").Insert(siafile)
	if err != nil {
		sf.logger.Println("ERROR: Could not save record into database.")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		sf.logger.Println("ERROR: Could not encode siafile")
		return
	}
}

//  GetAll returns all siafiles as a json list.
func (sf *Sunfish) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO Pagination and sorts
	var siafiles []Siafile

	// Select removes the content from query results use for not returning .sia
	err := sf.DB.C("siafiles").Find(bson.M{"listed": true}).Sort("-uploadedtime").All(&siafiles)
	if err != nil {
		sf.logger.Println("ERROR: Could not find all siafiles.")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafiles); err != nil {
		sf.logger.Println(err)
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
		sf.logger.Println("ERROR: Could not find Siafile: ", bson.ObjectIdHex(id))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		sf.logger.Println("ERROR: Could not JSON encode Siafile:", siafile)
		return
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
	err := sf.DB.C("siafiles").Find(bson.M{"tags": search, "listed": true}).All(&siafiles)

	if err != nil {
		sf.logger.Println("ERROR: searching Siafiles for tags: ", search)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafiles); err != nil {
		sf.logger.Println("ERROR: Could not encode Siafile", siafiles)
		return
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
		sf.logger.Println(err)
		return
	}
}
