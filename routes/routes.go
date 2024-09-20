package routes

import (
	"tienda-electronica/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	//rutas clientes

	router.HandleFunc("/clientes", controllers.GetClientes(db)).Methods("GET")
	router.HandleFunc("/cliente", controllers.CrearCliente(db)).Methods("POST")
	router.HandleFunc("/cliente/{id}", controllers.EliminarCliente).Methods("DELETE")
	router.HandleFunc("/cliente/{id}", controllers.ActualizarCliente).Methods("PUT")
	router.HandleFunc("/clientes/{cliente_id}/ordenes", controllers.ObtenerOrdenesDeCliente(db)).Methods("GET")

	//rutas productos

	router.HandleFunc("/productos", controllers.GetProductos(db)).Methods("GET")
	router.HandleFunc("/producto", controllers.CrearProducto(db)).Methods("POST")
	router.HandleFunc("/producto/{id}", controllers.EliminarProducto).Methods("DELETE")
	router.HandleFunc("/producto/{id}", controllers.ActualizarProducto).Methods("PUT")
	router.HandleFunc("/productos/{producto_id}/ventas", controllers.ObtenerTotalVentasProducto(db)).Methods("GET")
	router.HandleFunc("/productos/mas-vendidos", controllers.ObtenerProductosMasVendidos(db)).Methods("POST")

	//rutas ordenes

	router.HandleFunc("/ordenes", controllers.GetOrdenes(db)).Methods("GET")
	router.HandleFunc("/ordenes/{id}/detalles", controllers.GetOrdenConDetalles(db)).Methods("GET")
	router.HandleFunc("/orden", controllers.CrearOrdenConDetalles(db)).Methods("POST")
	router.HandleFunc("/orden/{id}", controllers.EliminarOrden).Methods("DELETE")
	router.HandleFunc("/orden/{id}", controllers.ActualizarOrden).Methods("PUT")

	return router
}
