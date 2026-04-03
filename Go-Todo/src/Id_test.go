package main

import (
	"testing"
)

func nuevoGen() *GenerarID {
	return &GenerarID{}
}

// --- Tests de ID ---

func TestIDsUnicosPorProducto(t *testing.T) {
	g := nuevoGen()
	laptop := NewElectronicoConGenerador("Laptop", 900.00, 3, "Dell", 12, g)
	remera := NewRopaConGenerador("Remera", 30.00, 5, "L", "Algodón", g)
	libro := NewLibroConGenerador("Go in Action", 50.00, 8, "Kennedy", "Programación", g)

	if laptop.GetID() == remera.GetID() || remera.GetID() == libro.GetID() {
		t.Error("cada producto debe tener un ID único")
	}
}
