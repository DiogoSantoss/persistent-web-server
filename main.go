// Package main initializes webserver on PORT
package main

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/router"

	"log"
	"net/http"
)

const PORT string = ":5000"

func main() {

	log.Printf("Running server on localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router.Router()))
}

// TODOS:
// - Change middleware function names to be more descriptive
// - Add real database
// - Add error handling
// - Format GET response to return data as json
// - Improve http response messages
// - Add documentation
// OPTIONAL:
// - Add tests
