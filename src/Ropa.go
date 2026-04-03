package main

import (
	"fmt"
)

type Ropa struct {
	ProductoBase
	talle    string
	material string
}

func NewRopa(nombre string, precio float64, stock int, talle, material string) *Ropa {
	return &Ropa{
		ProductoBase: nuevoProductoBase(nombre, precio, stock), //Llama a `nuevoProductoBase()` que internamente genera el ID automático y devuelve un `ProductoBase` listo
		talle:        talle,
		material:     material,
	}
}

func NewRopaConGenerador(nombre string, precio float64, stock int, talle, material string, g *GenerarID) *Ropa {
	return &Ropa{
		ProductoBase: nuevoProductoBaseConGenerador(nombre, precio, stock, g),
		talle:        talle,
		material:     material,
	}
}

func (r *Ropa) Descripcion() string {
	return fmt.Sprintf("[Ropa] #%d | %s | Talle: %s | Material: %s | $%.2f",
		r.id, r.nombre, r.talle, r.material, r.precio)
}
