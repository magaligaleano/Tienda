package models

import (
	"gorm.io/gorm"
)

type DetalleOrden struct {
	gorm.Model
	OrdenID        uint     `json:"orden_id"`
	ProductoID     uint     `json:"producto_id"`
	Cantidad       int      `json:"cantidad"`
	PrecioUnitario float64  `json:"precio_unitario"`
	Producto       Producto `json:"producto,omitempty"`
}
