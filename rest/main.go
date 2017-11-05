// Serve up the bets via a simple REST API,
// for me to use while I learn a bit of angular.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	port := os.Getenv("REST_PORT")
	fmt.Println("Listening on port", port)
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
