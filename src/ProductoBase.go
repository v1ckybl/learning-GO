package main

import (
	"fmt"
)

type ProductoBase struct { //clase abstracta que implementa la interfaz Producto, sirve como base para otros tipos de productos
	id     uint64
	nombre string
	precio float64
	stock  int
}

func nuevoProductoBase(nombre string, precio float64, stock int) ProductoBase {
	return ProductoBase{
		id:     generarID(), // ID asignado automáticamente al construir
		nombre: nombre,
		precio: precio,
		stock:  stock,
	}
}

func (p *ProductoBase) GetID() uint64 {
	return p.id
}

func (p *ProductoBase) GetNombre() string {
	return p.nombre
}

func (p *ProductoBase) GetPrecio() float64 {
	return p.precio
}

func (p *ProductoBase) GetStock() int {
	return p.stock
}

func (p *ProductoBase) EstaDisponible() bool {
	return p.stock > 0
}

func (p *ProductoBase) DescontarStock(cantidad int) error {
	if p.stock < cantidad {
		return fmt.Errorf("stock insuficiente: disponible %d, pedido %d", p.stock, cantidad)
	}
	p.stock -= cantidad
	return nil
}
