package main

type IProducto interface { //contrato que todo producto debe cumplir
	GetID() uint64
	GetNombre() string
	GetPrecio() float64
	GetStock() int
	EstaDisponible() bool
	DescontarStock(cantidad int) error
}
