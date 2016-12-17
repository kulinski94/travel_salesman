package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/api/cities/{count}",
		GenerateCities,
	},
	Route{
		"Path",
		"POST",
		"/api/path",
		FindPath,
	},
}
