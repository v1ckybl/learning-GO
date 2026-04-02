package main

import (
	"fmt"
)

type Libro struct {
	ProductoBase
	autor     string
	categoria string
}

func NewLibro(nombre string, precio float64, stock int, autor, categoria string) *Libro {
	return &Libro{
		ProductoBase: nuevoProductoBase(nombre, precio, stock),
		autor:        autor,
		categoria:    categoria,
	}
}

func (l *Libro) Descripcion() string {
	return fmt.Sprintf("[Libro] #%d | %s | Autor: %s | Categoría: %s | $%.2f",
		l.id, l.nombre, l.autor, l.categoria, l.precio)
}
