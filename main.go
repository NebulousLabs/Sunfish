package main

import (
	"flag"
	"net/http"
)

func main() {
	// Add flag for listening on a different port
	port := flag.String("port", "10181", "Port to listen on.")
	dbName := flag.String("dbName", "sunfish", "Mongo DB to use")
	logDir := flag.String("logDir", "logs", "Directory to put log file in.")
	flag.Parse()

	sf := NewSunfish(*logDir, *dbName)
	defer sf.Close()
	err := http.ListenAndServe(":"+*port, sf.Router)
	if err != nil {
		sf.logger.Fatalln("ERROR: Attempting to listen and serve on port: " + *port)
	}
}
