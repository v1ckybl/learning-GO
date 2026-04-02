package main

type ProductoBase struct {
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
