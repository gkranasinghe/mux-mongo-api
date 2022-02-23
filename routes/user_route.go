package routes

import (
    "mux-mongo-api/controllers" //add this
    "github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
    router.HandleFunc("/user", controllers.CreateUser()).Methods("POST") //add this
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET") //add this
	router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT") //add this
}