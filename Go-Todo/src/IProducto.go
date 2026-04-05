package main

type IProducto interface { //contrato que todo producto debe cumplir
	GetID() uint64
	GetNombre() string
	GetPrecio() float64
	GetStock() int
	EstaDisponible() bool
	Reservar(cantidad int) error
	LiberarReserva(cantidad int)
	ConfirmarCompra(cantidad int) error
}
