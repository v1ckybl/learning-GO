package main

type ItemCompra struct {
	producto IProducto // uso de INTERFAZ como tipo
	cantidad int
	subtotal float64
}

func newItem(p IProducto, cantidad int) ItemCompra {
	return ItemCompra{
		producto: p,
		cantidad: cantidad,
		subtotal: p.GetPrecio() * float64(cantidad),
	}
}
