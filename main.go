package main

import (
	"flag"
	"net/http"
)

func main() {
	// Add flag for listening on a different port
	port := flag.String("port", "10181", "Port to listen on.")
	flag.Parse()

	sf := NewSunfish()
	defer sf.Close()
	http.ListenAndServe(":"+*port, sf.Router)
}
