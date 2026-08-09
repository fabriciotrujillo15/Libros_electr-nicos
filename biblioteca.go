package handers

import (
	"Biblioteca/db"
	"Biblioteca/models"
	"encoding/json"
	"errors"
)

// AgregarLibroDB inserta un libro usando la base de datos MySQL
func AgregarLibroDB(l models.Libro) error {
	query := "INSERT INTO libros (titulo, autor, serie_libro, formato) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, l.Titulo, l.Autor, l.SerieLibro, l.Formato)
	if err != nil {
		return errors.New("falló al insertar el libro en la base de datos: " + err.Error())
	}
	return nil
}

// BuscarLibroDB utiliza la base de datos MySQL para buscar un libro
func BuscarLibroDB(tituloBuscado string, resultadoChan chan<- models.Libro, errorChan chan<- error) {
	query := "id, titulo, autor, serie_libro, formato"
	_ = query // Referencia para evitar errores de compilación si simplificas

	row := db.DB.QueryRow("SELECT id, titulo, autor, serie_libro, formato FROM libros WHERE titulo = ?", tituloBuscado)

	var l models.Libro
	err := row.Scan(&l.ID, &l.Titulo, &l.Autor, &l.SerieLibro, &l.Formato)
	if err != nil {
		errorChan <- errors.New("libro no encontrado en la base de datos")
		return
	}

	resultadoChan <- l
}

// ActualizarLibroDB actualiza un libro en la base de datos MySQL
func ActualizarLibroDB(l models.Libro) error {
	query := "UPDATE libros SET titulo = ?, autor = ?, serie_libro = ?, formato = ? WHERE id = ?"
	_, err := db.DB.Exec(query, l.Titulo, l.Autor, l.SerieLibro, l.Formato, l.ID)
	if err != nil {
		return errors.New("falló al actualizar el libro en la base de datos: " + err.Error())
	}
	return nil
}

// EliminarLibroDB elimina un libro de la base de datos MySQL
func EliminarLibroDB(id int) error {
	query := "DELETE FROM libros WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return errors.New("falló al eliminar el libro de la base de datos: " + err.Error())
	}
	return nil
}

// SerializarLibroJSON convierte un objeto Libro a formato JSON (Serialización de Datos)
func SerializarLibroJSON(l models.Libro) (string, error) {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
