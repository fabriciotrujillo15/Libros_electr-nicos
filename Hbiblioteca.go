package handers

import (
	"Biblioteca/db"
	"Biblioteca/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// funcion agregar en la base de datos
func AgregarLibroDB(l models.Libro) error {
	query := "INSERT INTO Libros (titulo, autor, serie_libro, formato, Usuarios_id) VALUES (@p1, @p2, @p3, @p4, @p5)"
	_, err := db.DB.Exec(query, l.Titulo, l.Autor, l.SerieLibro, l.Formato, "admin_default")
	return err
}
// funcion actualizar en la base de datos
func ActualizarLibroDB(l models.Libro) error {
	query := "UPDATE Libros SET titulo = @p1, autor = @p2, serie_libro = @p3, formato = @p4 WHERE id = @p5"
	_, err := db.DB.Exec(query, l.Titulo, l.Autor, l.SerieLibro, l.Formato, l.ID)
	return err
}
// funcion eliminar en la base de datos
func EliminarLibroDB(id int) error {
	query := "DELETE FROM Libros WHERE id = @p1"
	_, err := db.DB.Exec(query, id)
	return err
}

// Función para renderizar la página de inicio
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/base.html", "views/bibiliotecaD.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Pagina": "inicio",
	}
	tmpl.Execute(w, data)
}

// funcion para listar los libros en la base de datos y mostrarlos en la plantilla HTML
func ListarLibrosHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, titulo, autor, serie_libro, formato FROM Libros")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // cerrar 

	var libros []models.Libro
	for rows.Next() {
		var l models.Libro
		if err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.SerieLibro, &l.Formato); err != nil {
			http.Error(w, "Error al leer libros: "+err.Error(), http.StatusInternalServerError)
			return
		}
		libros = append(libros, l)
	}

	tmpl, err := template.ParseFiles("views/base.html", "views/libros.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Libros": libros,
		"Pagina": "libros",
	}
	tmpl.Execute(w, data)
}

// funcion para crear un libro en la base de datos y redirigir a la lista de libros
func CrearLibroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		l := models.Libro{
			Titulo:     r.FormValue("titulo"),
			Autor:      r.FormValue("autor"),
			SerieLibro: r.FormValue("serie_libro"),
			Formato:    r.FormValue("formato"),
		}
		AgregarLibroDB(l)
		http.Redirect(w, r, "/libros", http.StatusSeeOther)
	}
}

// funcion para actualizar un libro en la base de datos y redirigir a la lista de libros
func ActualizarLibroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var id int
		fmt.Sscanf(r.FormValue("id"), "%d", &id)

		l := models.Libro{
			ID:         id,
			Titulo:     r.FormValue("titulo"),
			Autor:      r.FormValue("autor"),
			SerieLibro: r.FormValue("serie_libro"),
			Formato:    r.FormValue("formato"),
		}
		ActualizarLibroDB(l)
		http.Redirect(w, r, "/libros", http.StatusSeeOther)
	}
}

// Funcion para eliminar un libro de la base de datos y redirigir a la lista de libros
func EliminarLibroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var id int
		fmt.Sscanf(r.FormValue("id"), "%d", &id)
		EliminarLibroDB(id)
		http.Redirect(w, r, "/libros", http.StatusSeeOther)
	}
}

// Funcion para api listar los libros en la base de datos y mostrarlos en formato JSON
func ApiListarLibrosHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, titulo, autor, serie_libro, formato FROM Libros")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Corregido aquí también

	var libros []models.Libro
	for rows.Next() {
		var l models.Libro
		rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.SerieLibro, &l.Formato)
		libros = append(libros, l)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}
