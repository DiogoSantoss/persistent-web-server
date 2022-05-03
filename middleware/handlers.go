// Package middleware provides handlers to deal with API requests
// and stores data in a database.
package middleware

import (
	"log"

	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/models"

	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

// Temporary "database", to be replaced by a real database
var db []models.Data
var counter int = 0

// Decoder to transform form data into struct
var decoder = schema.NewDecoder()


func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(db)
}

func PutDataPost(w http.ResponseWriter, r *http.Request) {

	var data models.Data

	err := r.ParseMultipartForm(1024) 
	if err != nil {
		log.Fatalf("Failed to parse form with error: %v", err)
	}

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		log.Fatalf("Failed to decode form with error: %v", err)
	}

	data.ID = int64(counter)
	counter++

	log.Printf("New entry %v added to database\n", data)
	db = append(db, data)
}

func PutDataGet(w http.ResponseWriter, r *http.Request) {

	var data models.Data
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		log.Fatalf("Failed to decode parameters with error: %v", err)
	}

	data.ID = int64(counter)
	counter++

	log.Printf("New entry %v added to database\n", data)
	db = append(db, data)
}