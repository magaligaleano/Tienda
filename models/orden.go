package models

import (
	"gorm.io/gorm"
)

type Orden struct {
	gorm.Model
	ClienteID uint           `json:"cliente_id"`
	Fecha     string         `json:"fecha"`
	Total     float64        `json:"total"`
	Detalles  []DetalleOrden `json:"detalles" gorm:"constraint:OnDelete:CASCADE;"`
}
