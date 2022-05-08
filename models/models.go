// Package models provides structs
package models

import (
	"github.com/jinzhu/gorm"
)

// Single data stored in the database
type Data struct {
	gorm.Model
	
	Latitude	int64		
	Longitude	int64		
	Time		string
	Speed		int64
}

// Aggregate data not stored in the database
type AggregateData struct {

	Latitude 	[]int64
	Longitude 	[]int64
	Time 		[]string
	Speed 		[]int64
}

// Response sent to client
type Response struct {
	Data 		*Data				`json:"data,omitempty"`
	ListData	*AggregateData  	`json:"listData,omitempty"`
    Message 	string 				`json:"message,omitempty"`
}