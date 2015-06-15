package main

import "time"

type Siafile struct {
	_id          string    `json:"id",bson:"_id, omitempty"`
	SiafilePath  string    `json:"siafilepath"`
	Hash         string    `json:"hash"`
	Title        string    `json:"title"`
	UploadedTime time.Time `json:"uploadedTime"`
	Tags         []string  `json:"tags"`
}
