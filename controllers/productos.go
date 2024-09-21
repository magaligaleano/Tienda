package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tienda-electronica/config"
	"tienda-electronica/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CrearProducto(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var producto models.Producto
		if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if result := db.Create(&producto); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(producto)
	}
}
func GetProductos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productos []models.Producto
		db.Find(&productos)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)
	}
}

func EliminarProducto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := config.SetupDB()

	var producto models.Producto
	if err := db.First(&producto, id).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	if err := db.Unscoped().Delete(&producto).Error; err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := config.SetupDB()

	var producto models.Producto
	if err := db.First(&producto, id).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Save(&producto).Error; err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(producto)
}
func ObtenerTotalVentasProducto(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productoID, err := strconv.Atoi(vars["producto_id"])
		if err != nil {
			http.Error(w, "ID de producto inválido", http.StatusBadRequest)
			return
		}

		var totalVentas float64
		err = db.Model(&models.DetalleOrden{}).Select("SUM(cantidad * precio_unitario)").Where("producto_id = ?", productoID).Scan(&totalVentas).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"total_ventas": totalVentas})
	}
}
func ObtenerProductosMasVendidos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			FechaInicio string `json:"fecha_inicio"`
			FechaFin    string `json:"fecha_fin"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
			return
		}

		var productosVendidos []struct {
			ProductoID   uint   `json:"producto_id"`
			Nombre       string `json:"nombre"`
			TotalVendido int64  `json:"total_vendido"`
		}

		err := db.Table("detalle_ordens").
			Select("producto_id, productos.nombre, SUM(detalle_ordens.cantidad) as total_vendido").
			Joins("join productos on productos.id = detalle_ordens.producto_id").
			Joins("join ordens on ordens.id = detalle_ordens.orden_id").
			Where("ordens.fecha BETWEEN ? AND ?", params.FechaInicio, params.FechaFin).
			Group("producto_id, productos.nombre").
			Order("total_vendido DESC").
			Scan(&productosVendidos).Error

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productosVendidos)
	}
}
