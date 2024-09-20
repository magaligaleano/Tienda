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

func CrearCliente(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cliente models.Cliente
		if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if result := db.Create(&cliente); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(cliente)
	}
}
func GetClientes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientes []models.Cliente
		db.Find(&clientes)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clientes)
	}
}
func EliminarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := config.SetupDB()

	var cliente models.Cliente
	if err := db.First(&cliente, id).Error; err != nil {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	if err := db.Unscoped().Delete(&cliente).Error; err != nil {
		http.Error(w, "Error al eliminar el cliente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := config.SetupDB()

	var cliente models.Cliente
	if err := db.First(&cliente, id).Error; err != nil {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Save(&cliente).Error; err != nil {
		http.Error(w, "Error al actualizar el cliente", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cliente)
}
func ObtenerOrdenesDeCliente(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		clienteID, err := strconv.Atoi(vars["cliente_id"])
		if err != nil {
			http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
			return
		}

		var ordenes []models.Orden
		if err := db.Preload("DetalleOrden.Producto").Where("cliente_id = ?", clienteID).Find(&ordenes).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ordenes)
	}
}
