package main

import "fmt"

var siafiles Siafiles

// Stand in replacement for a database. will be replaced w/ a database
func AddSiafile(s Siafile) Siafile {
	siafiles = append(siafiles, s)
	fmt.Printf("%+v\n", siafiles)
	return s
}
