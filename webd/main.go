// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Header struct {
	Title string
	Week  string
}

type RowData struct {
	Headline    string
	Predictions []string
}

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		fmt.Println("Error renderTemplate ParseFiles: ", err.Error())
	}
	err = t.Execute(w, d)
	if err != nil {
		fmt.Println("Error renderTemplate Execute: ", err.Error())
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	d := &Header{
		Title: "NFL Betting Results",
		Week:  "1",
	}
	renderTemplate(w, "head", d)

	users := []string{"User1", "User2"}
	renderTemplate(w, "table_headings", users)

	pr := []string{"CHI (+2)", "CIN (-5)"}
	rd := &RowData{
		Headline:    "CHI 12 - 16 CIN",
		Predictions: pr,
	}
	renderTemplate(w, "row", rd)

	renderTemplate(w, "foot", "")
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
