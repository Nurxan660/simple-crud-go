package routes

import (
	"testGo/pkg/controllers"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) *mux.Router {

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}
