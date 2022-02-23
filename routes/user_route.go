package routes

import (
    "mux-mongo-api/controllers" //add this
    "github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
    router.HandleFunc("/user", controllers.CreateUser()).Methods("POST") //add this
}