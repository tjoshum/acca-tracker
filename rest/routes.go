package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GET",
		"/",
		"Index",
		Index,
	},
	Route{
		"GET",
		"/week/{weekNumber}",
		"getWeekBets",
		WeekNumber,
	},
}
