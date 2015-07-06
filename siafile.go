package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Siafile struct {
	Id           bson.ObjectId `bson:"_id,omitempty", json:"_id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Filename     string        `json:"filename"`
	FileData     string        `json:"fileData"`
	UploadedTime time.Time     `json:"uploadedTime"`
	Tags         []string      `json:"tags"`
}
