package main

import (
	"Biblioteca/db"
	"Biblioteca/handers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar .env")
	}

	// Inicializar conexión a la BD global
	db.InitDB()

	// Configurar Enrutador con Gorilla Mux
	r := mux.NewRouter()

	// 1. Rutas Web (Templates HTML para los usuarios)
	r.HandleFunc("/", handers.IndexHandler).Methods("GET")
	r.HandleFunc("/libros", handers.ListarLibrosHandler).Methods("GET")
	r.HandleFunc("/libros/crear", handers.CrearLibroHandler).Methods("POST")
	r.HandleFunc("/libros/eliminar", handers.EliminarLibroHandler).Methods("POST")
	r.HandleFunc("/libros/actualizar", handers.ActualizarLibroHandler).Methods("POST")

	// 2. Rutas de API REST (Para conexión con otros sistemas y verificación en navegador)
	r.HandleFunc("/api/libros", handers.ApiListarLibrosHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // El puerto que usa tu docente en la captura
	}

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
