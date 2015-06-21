package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Siafile struct {
	Id           bson.ObjectId `bson:"_id,omitempty", json:"_id"`
	Filename     string        `json:"filename"`
	Description  string        `json:"description"`
	Content      string        `json:"fileData"`
	Title        string        `json:"title"`
	UploadedTime time.Time     `json:"uploadedTime"`
	Tags         []string      `json:"tags"`
}
