package main

import "testing"

func TestIDsUnicosPorProducto(t *testing.T) {
	laptop := NewElectronico("Laptop", 900.00, 3, "Dell", 12)
	remera := NewRopa("Remera", 30.00, 5, "L", "Algodón")
	libro := NewLibro("Go in Action", 50.00, 8, "Kennedy", "Programación")
 
	if laptop.GetID() == remera.GetID() || remera.GetID() == libro.GetID() {
		t.Error("cada producto debe tener un ID único")
	}
}






/*func TestProducto(t *testing.T) {
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
}//*
