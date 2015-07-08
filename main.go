package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Add flag for listening on a different port
	port := flag.String("port", "10181", "Port to listen on.")
	logDir := flag.String("logDir", "logs", "Directory to put log file in.")
	flag.Parse()

	sf := NewSunfish(*logDir)
	defer sf.Close()
	err := http.ListenAndServe(":"+*port, sf.Router)
	if err != nil {
		fmt.Println("Error attempting to listen and serve on port: " + *port)
		os.Exit(1)
	}
}
