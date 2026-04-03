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
