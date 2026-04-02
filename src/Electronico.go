package main

import (
	"fmt"
)

type Electronico struct {
	ProductoBase
	marca    string
	garantia int // meses
}

func NewElectronico(nombre string, precio float64, stock int, marca string, garantia int) *Electronico {
	return &Electronico{
		ProductoBase:  nuevoProductoBase(nombre, precio, stock),
		marca:    marca,
		garantia: garantia,
	}
}

func (e *Electronico) Descripcion() string {
	return fmt.Sprintf("[Electrónico] #%d | %s | Marca: %s | Garantía: %d meses | $%.2f",
		e.id, e.nombre, e.marca, e.garantia, e.precio)
} //dudoso a ver che