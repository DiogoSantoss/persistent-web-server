// Package router implements the Router function that
// creates a mux.Router and registers routes
package router

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/get_data", middleware.GetData).Methods("GET")
	router.HandleFunc("/put_data", middleware.PutDataPost).Methods("POST")
	router.HandleFunc("/put_data", middleware.PutDataGet).Methods("GET")

	return router
}