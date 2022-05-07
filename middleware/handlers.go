// Package middleware provides handlers to deal with API requests
// and stores data in a database.
package middleware

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/models"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

// Max size of form
const MAX_FORM_SIZE = 1024

// Decoder to transform form data into struct
var decoder = schema.NewDecoder()

// get_data handler returns all data from database
func GetData(w http.ResponseWriter, r *http.Request) {

	var data = getDataDatabase()

	log.Println(data)

	// Concat all data into a single struct
	var allData models.AggregateData
	for _, d := range data {
		allData.Latitude = append(allData.Latitude, d.Latitude)
		allData.Longitude = append(allData.Longitude, d.Longitude)
		allData.Time = append(allData.Time, d.Time)
		allData.Speed = append(allData.Speed, d.Speed)
	}

	res := models.Response {
		ListData: &allData,
		Message: "Data retrieved successfully",
	}

	// send response
	json.NewEncoder(w).Encode(res)
	
	// log response
	log.Printf("Sent %v entries from database\n", len(data))
}

// put_data handler receives data from POST request
func PutDataPost(w http.ResponseWriter, r *http.Request) {

	var data models.Data

	err := r.ParseForm() 
	if err != nil {
		log.Fatalf("Failed to parse form with error: %v", err)
	}

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		log.Fatalf("Failed to decode form with error: %v", err)
	}

	addDataDatabase(data)

	res := models.Response {
		Data: &data,
		Message: "Data added successfully",
	}

	// send respsonse
	json.NewEncoder(w).Encode(&res)

	// log response
	log.Printf("New entry %v added to database\n", data)
}

// put_data handler receives data from GET request
func PutDataGet(w http.ResponseWriter, r *http.Request) {

	var data models.Data

	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		log.Fatalf("Failed to decode parameters with error: %v", err)
	}

	addDataDatabase(data)

	res := models.Response {
		Data: &data,
		Message: "Data added successfully",
	}

	// send respsonse
	json.NewEncoder(w).Encode(&res)

	// log response
	log.Printf("New entry %v added to database\n", data)
}