package main

type Producto interface {
	ID int
	Nombre string
	Precio float64
	Stock int
	Categoria string
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

