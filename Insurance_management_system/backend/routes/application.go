package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

// RegisterApplicationRoutes sets up routes for application-related operations
func RegisterApplicationRoutes(router *mux.Router) {
	// Create a new application
	router.HandleFunc("/applications", controllers.CreateApplication).Methods("POST")

	// Get all applications for a specific user (using user_id)
	router.HandleFunc("/users/{user_id}/applications", controllers.GetApplicationsByUserID).Methods("GET")

	// Get profile by username and email
	router.HandleFunc("/users/profile", controllers.GetUserProfile).Methods("GET")

	// Approve or reject applications
	router.HandleFunc("/applications/approve/{id}", controllers.ApproveApplication).Methods("POST")
	router.HandleFunc("/applications/reject/{id}", controllers.RejectApplication).Methods("POST")

	// PUT method to update an application**
	router.HandleFunc("/applications/{application_id}", controllers.UpdateApplication).Methods("PUT")

	// Delete an application by its ID
	router.HandleFunc("/applications/{application_id}", controllers.DeleteApplication).Methods("DELETE")

	// Get all applications
	router.HandleFunc("/applications", controllers.GetAllApplications).Methods("GET")

	// Get Passport Photo
	router.HandleFunc("/getPassportPhoto/{user_id}", controllers.GetPassportPhoto).Methods("GET")

	// Get ID Proof
	router.HandleFunc("/getIdProof/{user_id}", controllers.GetIdProof).Methods("GET")

	// Get applications by name
	router.HandleFunc("/applications/{username}", controllers.GetApplicationsByName).Methods("GET")
}
