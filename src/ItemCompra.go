package main

type ItemCompra struct {
	Producto Producto
	Cantidad int
}

func (i ItemCompra) PrecioPorItem() float64 {
	i.Producto.Precio * float64(i.Cantidad)
}
