package main

import (
	"log"
	"net/http"
	"tienda-electronica/config"
	"tienda-electronica/models"
	"tienda-electronica/routes"

	"github.com/rs/cors"
)

func main() {

	config.LoadEnv()

	db := config.SetupDB()

	db.AutoMigrate(&models.Cliente{}, &models.Producto{}, &models.Orden{}, &models.DetalleOrden{})

	router := routes.SetupRouter(db)

	corsHandler := cors.Default().Handler(router)

	log.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
