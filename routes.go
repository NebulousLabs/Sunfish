package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func newRouter(sf *Sunfish) *mux.Router {
	var routes = []Route{
		Route{
			Name:        "Index",
			Method:      "POST",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.AddFile,
		},
		Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.GetAll,
		},
		Route{
			Name:        "Get Siafile",
			Method:      "GET",
			Pattern:     "/api/siafile/{id}",
			HandlerFunc: sf.GetFile,
		},
		Route{
			Name:        "Search",
			Method:      "GET",
			Pattern:     "/api/siafile/search/",
			HandlerFunc: sf.SearchFile,
		},
		Route{
			Name:        "Delete",
			Method:      "DELETE",
			Pattern:     "/api/siafile/",
			HandlerFunc: sf.DeleteFile,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Serve static folder at root
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return router
}
