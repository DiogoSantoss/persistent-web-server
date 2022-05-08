// Package main initializes webserver on PORT
package main

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/router"
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/middleware"

	"log"
	"net/http"
)

// Web server port
const PORT string = ":5000"

func main() {

	// Initialize database
	db := middleware.CreateConnection()

	// When the application is closed, close the database
	defer db.Close()

	// Create a new router with routes
	r := router.CreateRouter()

	log.Printf("Running server on localhost%s\n", PORT)

	// Listens for TCP connections on PORT
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("Failed to start server with error: %v\n",err)
	}
}

// TODOS:
// - When printing Data, gorm.Model appears null (ideally gorms.Model params dont show)
// - Add documentation
// OPTIONAL:
// - Add tests