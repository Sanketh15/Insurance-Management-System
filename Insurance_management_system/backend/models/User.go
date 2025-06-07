package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"` // Explicitly set the primary key
	Role     string `json:"role"`                 // Can be "admin" or "staff"
	Username string `json:"username"`             // Unique identifier for login
	Password string `json:"password"`             // Password for authentication
	Email    string `json:"email"`                // Email for the user
}
