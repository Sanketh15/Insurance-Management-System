package routes

import (
	"backend/controllers"
	"github.com/gorilla/mux"
)

// SetupUserRoutes defines the routes related to users
func SetupUserRoutes(r *mux.Router) {
	r.HandleFunc("/userregister", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/userDetails", controllers.GetUserDetails).Methods("GET") // Updated line
}
