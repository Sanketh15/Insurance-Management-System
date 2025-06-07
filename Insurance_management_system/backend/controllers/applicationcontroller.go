package controllers

import (
	"backend/config"
	"backend/models"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// GetApplicationsByUserID retrieves all applications for a specific user
func GetApplicationsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var applications []models.Application
	if err := config.DB.Where("user_id = ?", userID).Find(&applications).Error; err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}

// ApproveApplication sets the Approved flag to true and clears the Rejected flag.
func ApproveApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"] // Use "id" consistently
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid application id", http.StatusBadRequest)
		return
	}

	var application models.Application
	if err := config.DB.First(&application, id).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	// Expect a JSON body: { "approved": true }
	var requestBody struct {
		Approved bool `json:"approved"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	if requestBody.Approved {
		application.Approved = true
		application.Rejected = false
		application.Reason = ""
	} else {
		// If the client sends approved: false, then we simply mark it not approved.
		application.Approved = false
	}

	if err := config.DB.Save(&application).Error; err != nil {
		http.Error(w, "Failed to update application status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}

// RejectApplication sets the Rejected flag to true and clears the Approved flag.
func RejectApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"] // Use "id" consistently
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid application id", http.StatusBadRequest)
		return
	}

	var application models.Application
	if err := config.DB.First(&application, id).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	// Expect a JSON body: { "rejection_reason": "reason text" }
	var requestBody struct {
		RejectionReason string `json:"rejection_reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	application.Approved = false
	application.Rejected = true
	application.Reason = requestBody.RejectionReason

	if err := config.DB.Save(&application).Error; err != nil {
		http.Error(w, "Failed to reject application", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}

// DeleteApplication deletes an application by its ID
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	var application models.Application
	if err := config.DB.First(&application, appID).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&application).Error; err != nil {
		http.Error(w, "Failed to delete application", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Application deleted successfully"))
}

// CreateApplication handles the submission of a new application
func CreateApplication(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	pan := r.FormValue("pan")
	salaryStr := r.FormValue("salary")
	insurancePlan := r.FormValue("insurancePlan")

	if name == "" || email == "" || phone == "" || pan == "" || salaryStr == "" || insurancePlan == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		http.Error(w, "Invalid salary value", http.StatusBadRequest)
		return
	}

	passportPhoto, photoHeader, err := r.FormFile("passportPhoto")
	if err != nil {
		http.Error(w, "Passport photo is required", http.StatusBadRequest)
		return
	}
	defer passportPhoto.Close()

	idProof, idProofHeader, err := r.FormFile("idProof")
	if err != nil {
		http.Error(w, "ID proof is required", http.StatusBadRequest)
		return
	}
	defer idProof.Close()

	passportPhotoPath, err := saveUploadedFile(passportPhoto, photoHeader.Filename, "uploads/passportPhotos")
	if err != nil {
		http.Error(w, "Failed to save passport photo", http.StatusInternalServerError)
		return
	}

	idProofPath, err := saveUploadedFile(idProof, idProofHeader.Filename, "uploads/idProofs")
	if err != nil {
		http.Error(w, "Failed to save ID proof", http.StatusInternalServerError)
		return
	}

	application := models.Application{
		Name:          name,
		Email:         email,
		Phone:         phone,
		PAN:           pan,
		Salary:        salary,
		InsurancePlan: insurancePlan,
		PassportPhoto: passportPhotoPath,
		IDProof:       idProofPath,
		Approved:      false,
	}

	if err := config.DB.Create(&application).Error; err != nil {
		http.Error(w, "Failed to create application", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(application)
}

func saveUploadedFile(file multipart.File, filename, uploadPath string) (string, error) {
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	dst, err := os.Create(filepath.Join(uploadPath, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filepath.Join(uploadPath, filename), nil
}

// GetPassportPhoto serves the passport photo of an application
func GetPassportPhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	var application models.Application
	if err := config.DB.First(&application, appID).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	photoPath := application.PassportPhoto
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		http.Error(w, "Passport photo not found", http.StatusNotFound)
		return
	}

	// Serve the photo file
	http.ServeFile(w, r, photoPath)
}

// GetIdProof serves the ID proof (could be image or PDF)
func GetIdProof(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	var application models.Application
	if err := config.DB.First(&application, appID).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	idProofPath := application.IDProof
	if _, err := os.Stat(idProofPath); os.IsNotExist(err) {
		http.Error(w, "ID proof not found", http.StatusNotFound)
		return
	}

	// Determine file type and set the Content-Type header
	ext := filepath.Ext(idProofPath)
	switch ext {
	case ".pdf":
		w.Header().Set("Content-Type", "application/pdf")
	case ".jpg", ".jpeg", ".png", ".gif":
		w.Header().Set("Content-Type", "image/jpeg")
	}

	// Serve the ID proof file (image or PDF)
	http.ServeFile(w, r, idProofPath)
}

// UpdateApplication handles updates to an existing application
func UpdateApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	var application models.Application
	if err := config.DB.First(&application, appID).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	if err := config.DB.Save(&application).Error; err != nil {
		http.Error(w, "Failed to update application", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(application)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	// Validate that both parameters are provided
	if username == "" || email == "" {
		http.Error(w, "username and email are required", http.StatusBadRequest)
		return
	}

	// Define a variable to hold the application record.
	// Note: In your applications table, the "name" field holds the username.
	var application models.Application

	// Use GORM to filter the applications table by name (username) and email.
	if err := config.DB.Where("name = ? AND email = ?", username, email).First(&application).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the matching application as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(application)
}
func GetAllApplications(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	var applications []models.Application
	query := config.DB

	// Filter by username and email
	if username != "" && email != "" {
		query = query.Where("name = ? AND email = ?", username, email)
	}

	if err := query.Order("created_at DESC").Find(&applications).Error; err != nil {
		http.Error(w, "Error fetching applications", http.StatusInternalServerError)
		return
	}

	// If no approved applications, include the latest unapproved one (optional)
	if len(applications) == 0 {
		if err := query.Order("created_at DESC").First(&applications).Error; err != nil {
			http.Error(w, "No applications found", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}

// GetApplicationByID retrieves an application by its ID.
func GetApplicationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["application_id"]

	var application models.Application
	if err := config.DB.First(&application, appID).Error; err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}

// GetApplicationsByName retrieves all applications for a specific user by name
func GetApplicationsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["username"] // Extract "name" from URL parameters

	var applications []models.Application
	// Query applications where the 'name' field matches the provided name
	if err := config.DB.Where("LOWER(name) = LOWER(?)", name).Find(&applications).Error; err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}
