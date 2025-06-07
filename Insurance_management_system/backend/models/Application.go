package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Name          string  `json:"name" validate:"required"`
	Email         string  `json:"email" validate:"required,email"`
	Phone         string  `json:"phone" validate:"required"`
	PAN           string  `json:"pan" validate:"required"`
	Salary        float64 `json:"salary" validate:"required"`
	InsurancePlan string  `json:"insurance_plan" validate:"required"`
	PassportPhoto string  `json:"passport_photo"`
	IDProof       string  `json:"id_proof"`
	Approved      bool    `json:"approved" gorm:"default:false"`
	Reason        string  `json:"Reason" gorm:"default:''"`
	Rejected      bool    `json:"rejected"`
}
