package main

import "time"

type Siafile struct {
	siafilePath  string    `json:"siafilepath"`
	Title        string    `json:"title"`
	UploadedTime time.Time `json:"uploadedTime"`
	Tags         []string  `json:"tags"`
}

type Siafiles []Siafile
