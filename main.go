package main

import (
	"net/http"
)

func main() {
	sf := NewSunfish()
	defer sf.Close()
	http.ListenAndServe(":10181", sf.Router)
}
