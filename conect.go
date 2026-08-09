// conectar base de datos
//conect.go

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConectarBD() {
	godotenv.Load()
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error al conectar: ", err)
	}
	fmt.Println("Conectado a SQL Server 2022")
}
