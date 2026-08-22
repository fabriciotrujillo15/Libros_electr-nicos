package models

import "errors"

// Interfaz para cumplir con el polimorfismo requerido
type Descargable interface {
	Descargar() (string, error)
}

// Estructura base Usuario (para herencia simulada en Go)
type Usuario struct {
	Nombre     string
	Email      string
	Cedula     string
	contrasena string // Encapsulado (minúscula)
}

// Getters y Setters para Encapsulación
func (u *Usuario) SetContrasena(pass string) error {
	if len(pass) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}
	u.contrasena = pass
	return nil
}

func (u *Usuario) GetContrasena() string {
	return u.contrasena
}

// Estructuras Administrador y Lector heredan lógicamente de Usuario
type Administrador struct {
	Usuario
	NivelAcceso int
}

type Lector struct {
	Usuario
	Suscripto bool
}

// Estructura Libro con anotaciones para Serialización JSON
type Libro struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	SerieLibro string `json:"serie_libro"`
	Formato    string `json:"formato"`
}

// Implementación de la interfaz Descargable (Polimorfismo)
func (l Libro) Descargar() (string, error) {
	if l.Formato == "" {
		return "", errors.New("formato de libro no válido para descarga")
	}
	return "Descargando el libro: " + l.Titulo + " en formato " + l.Formato, nil
}

//getters y setters para la estructura Libro
func (l *Libro) SetTitulo(titulo string) {
	l.Titulo = titulo
}

func (l *Libro) GetTitulo() string {
	return l.Titulo
}

func (l *Libro) SetAutor(autor string) {
	l.Autor = autor
}

func (l *Libro) GetAutor() string {
	return l.Autor
}

func (l *Libro) SetSerieLibro(serie string) {
	l.SerieLibro = serie
}

func (l *Libro) GetSerieLibro() string {
	return l.SerieLibro
}

func (l *Libro) SetFormato(formato string) {
	l.Formato = formato
}

func (l *Libro) GetFormato() string {
	return l.Formato
}

// Función para agregar un nuevo libro
func AgregarLibro(titulo string, autor string, serie string, formato string) *Libro {
	return &Libro{
		Titulo:     titulo,
		Autor:      autor,
		SerieLibro: serie,
		Formato:    formato,
	}
}

// Función para actualizar un libro existente
func (l *Libro) ActualizarLibro(titulo string, autor string, serie string, formato string) {
	l.Titulo = titulo
	l.Autor = autor
	l.SerieLibro = serie
	l.Formato = formato
}

// Función para eliminar un libro
func (l *Libro) EliminarLibro() {
	l.Titulo = ""
	l.Autor = ""
	l.SerieLibro = ""
	l.Formato = ""
}

// Función para mostrar información del libro
func (l *Libro) MostrarInfo() string {
	return "Título: " + l.Titulo + ", Autor: " + l.Autor + ", Serie: " + l.SerieLibro + ", Formato: " + l.Formato
}

// funcion para buscar un libro por título en una lista de libros
func BuscarLibroPorTitulo(libros []Libro, titulo string) (*Libro, error) {
	for _, libro := range libros {
		if libro.Titulo == titulo {
			return &libro, nil
		}
	}
	return nil, errors.New("libro no encontrado")
}

// funcion para eliminar un libro por título en una lista de libros
func EliminarLibroPorTitulo(libros *[]Libro, titulo string) error {
	for i, libro := range *libros {
		if libro.Titulo == titulo {
			*libros = append((*libros)[:i], (*libros)[i+1:]...)
			return nil
		}
	}
	return errors.New("libro no encontrado")
}
