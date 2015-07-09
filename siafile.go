package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Siafile is a struct containing all the Siafile info and metadata associated
// with and upload
type Siafile struct {
	Id           bson.ObjectId `bson:"_id,omitempty", json:"_id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Filename     string        `json:"filename"`
	Ascii        string        `json:"ascii"`
	UploadedTime time.Time     `json:"uploadedTime"`
	Tags         []string      `json:"tags"`
}
