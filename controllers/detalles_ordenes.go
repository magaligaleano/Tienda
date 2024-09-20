package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tienda-electronica/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetOrdenConDetalles(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orden models.Orden
		id := mux.Vars(r)["id"]

		ordenID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		if err := db.Preload("Detalles").First(&orden, ordenID).Error; err != nil {
			http.Error(w, "Orden no encontrada", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orden)
	}
}

func CrearOrdenConDetalles(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orden models.Orden

		if err := json.NewDecoder(r.Body).Decode(&orden); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&orden).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orden)
	}
}
