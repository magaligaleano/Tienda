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

func CrearOrden(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var orden models.Orden
		if err := json.NewDecoder(r.Body).Decode(&orden); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if result := db.Create(&orden); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orden)
	}
}
func GetOrdenes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ordenes []models.Orden
		db.Find(&ordenes)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ordenes)
	}
}
func EliminarOrden(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ordenID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := config.SetupDB()

	var orden models.Orden

	if err := db.First(&orden, ordenID).Error; err != nil {
		http.Error(w, "Orden no encontrada", http.StatusNotFound)
		return
	}

	if err := db.Unscoped().Delete(&orden).Error; err != nil {
		http.Error(w, "Error al eliminar la orden", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ActualizarOrden(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ordenID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	db := config.SetupDB()

	var orden models.Orden
	if err := db.First(&orden, ordenID).Error; err != nil {
		http.Error(w, "Orden no encontrada", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&orden); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Save(&orden).Error; err != nil {
		http.Error(w, "Error al actualizar la orden", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orden)
}
