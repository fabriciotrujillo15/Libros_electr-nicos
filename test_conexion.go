package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// 1. Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// 2. Construir la cadena de conexión
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// 3. Intentar conectar
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Error en la configuración de conexión: ", err)
	}
	defer db.Close()

	// 4. Probar si la base de datos responde
	err = db.Ping()
	if err != nil {
		log.Fatal("¡Error! No se pudo conectar a la base de datos: ", err)
	}

	fmt.Println("¡Conexión exitosa con Biblioteca_DB!")
}
