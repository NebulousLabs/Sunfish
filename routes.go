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
			"Index",
			"POST",
			"/api/siafile/",
			sf.AddFile,
		},
		Route{
			"Index",
			"GET",
			"/api/siafile/",
			sf.GetAll,
		},
		Route{
			"Get Siafile",
			"GET",
			"/api/siafile/{hash}",
			sf.GetFile,
		},
		Route{
			"Search",
			"GET",
			"/api/siafile/search/",
			sf.SearchFile,
		},
		Route{
			"Delete",
			"DELETE",
			"/api/siafile/",
			sf.DeleteFile,
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
