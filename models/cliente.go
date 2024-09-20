package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}
