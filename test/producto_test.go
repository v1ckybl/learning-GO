package main

import (
	"testing"
)

func TestProducto(t *testing.T) {
	p := Producto{
		Nombre: "pan",
		Precio: 200.0,
	}

	if p.Nombre != "pan" {
		t.Errorf("Nombre esperado 'pan', pero obtuve %v", p.Nombre)
	}

	if p.Precio != 200.0 {
		t.Errorf("Precio esperado 200.0, pero obtuve %f", p.Precio)
	}
}
