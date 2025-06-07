package routes

import (
	"backend/controllers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	
	r.HandleFunc("/auth/logout", controllers.Logout).Methods("POST")
}
