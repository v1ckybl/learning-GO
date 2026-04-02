package main

import (
	"fmt"
	"sync/atomic"
)

type IProducto interface { //contrato que todo producto debe cumplir
	GetID() uint64
	GetNombre() string
	GetPrecio() float64
	GetStock() int
	EstaDisponible() bool
	DescontarStock(cantidad int) error
}

/*func (p Producto) HayStock(cantidad int) bool{
	return p.Stock >= cantidad
}

func (p *Producto) ReducirStock(cantidad int) {
	if p.HayStock(cantidad) {
		p.Stock -= cantidad
		return true
	}
	return false
}/*

