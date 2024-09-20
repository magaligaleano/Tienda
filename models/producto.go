package models

import (
	"gorm.io/gorm"
)

// Producto representa un producto en la tienda.
type Producto struct {
	gorm.Model
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}
