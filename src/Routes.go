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
		"/api/question",
		getAllQuestions,
	},
	Route{
		"TodoIndex",
		"POST",
		"/api/question",
		addQuestion,
	},
	Route{
		"TodoShow",
		"DELETE",
		"/api/question/{questionId}",
		deleteQuestion,
	},
	Route{
		"Save pdf",
		"GET",
		"/api/quiz/pdf",
		savePdf,
	},
}
