package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Each Route object defines how a url is called and its handler function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// newRouter creates a router for the sunfish object with all the routes defined.
func newRouter(sf *Sunfish) *mux.Router {
	// Initializes all routes and their handlers and returns a mux
	var routes = []Route{
		// Handles a post request sending a new file
		Route{
			Name:        "Index",
			Method:      "POST",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.AddFile,
		},
		// Return all files in the database
		Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.GetAll,
		},
		// Search files using a '/?query=QUERY' in URL
		Route{
			Name:        "Search",
			Method:      "GET",
			Pattern:     "/api/siafile/search/",
			HandlerFunc: sf.SearchFile,
		},
		// Get a file by file id
		Route{
			Name:        "Get Siafile",
			Method:      "GET",
			Pattern:     "/api/siafile/{id}",
			HandlerFunc: sf.GetFile,
		},
		// Delete a siafile. Needs to be an authorized user
		Route{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.DeleteFile,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// Serve static folder at root
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return router
}
