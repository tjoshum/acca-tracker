package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bingo!")
}

func WeekNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekNumber := vars["weekNumber"]
	fmt.Fprintln(w, "Should fetch week:", weekNumber)
}
