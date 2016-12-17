package main

import (
	"net/http"
)

// Route - endpoint structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - list of Routre endpoint structure
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
		"/api/path/{algorithm}",
		FindPath,
	},
}
