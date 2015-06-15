package main

import (
	"encoding/json"
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

func writeUnmarshalError(w http.ResponseWriter) {
	var err error
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
}

func (sf *Sunfish) AddFile(w http.ResponseWriter, r *http.Request) {
	// Handles a Post Request for a Sia file and saves it to the DB
	var siafile Siafile
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &siafile); err != nil {
		writeUnmarshalError(w)
	}

	// Handle saving the Siafile to the db and file system
	siafile.UploadedTime = time.Now()
	err = sf.DB.C("siafiles").Insert(siafile)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		panic(err)
	}
}

func (sf *Sunfish) GetAll(w http.ResponseWriter, r *http.Request) {
	// Takes the hash of a Siafile in the URL and returns the Siafile in JSON
	var siafiles []Siafile

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

func (sf *Sunfish) GetFile(w http.ResponseWriter, r *http.Request) {
	// Takes the hash of a Siafile in the URL and returns the Siafile in JSON
	// TODO get hash from url
	//var hash string
	var siafile Siafile

	// TODO get file from DB and encode response
	// siafile = Db.C("siafiles").findOne({'hash':hash})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafile); err != nil {
		panic(err)
	}
}

func (sf *Sunfish) SearchFile(w http.ResponseWriter, r *http.Request) {
	// Takes a query parameter and searches the DB
	// TODO get query from URL
	// var query string
	var siafiles []Siafile

	// TODO Search database for query look up how to search mongo efficiently
	// siafiles = sf.DB.C("siafiles").Select(bson.M{tags: query})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(siafiles); err != nil {
		panic(err)
	}
}

func (sf *Sunfish) DeleteFile(w http.ResponseWriter, r *http.Request) {
	// Takes the hash of a Siafile in the URL and removes it from DB
	// TODO get hash from url
	var hash string

	// TODO get file from DB and encode response
	// Db.C("siafiles").remove({'hash': hash})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hash); err != nil {
		panic(err)
	}
}
