package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Error al abrir la BD: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la BD: ", err)
	}
	fmt.Println("Conexión a SQL Server establecida con éxito.")
}
