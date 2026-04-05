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
		t.Error("cada producto debe tener un ID unico")
	}
}

func TestStock(t *testing.T) {
	p := NewLibro("Las cronicas de narnia", 1000.05, 7, "C.S. Lewis", "Fantasía")
	_ = p.Reservar(2)
	// El stock original no deberia cambiar
	if p.GetStock() != 7 {
		t.Errorf("El stock original no deberia cambiar aun y se obtuvo %d", p.GetStockDisponible())
	}
	// El stock disponible si deberia cambiar
	if p.GetStockDisponible() != 5 {
		t.Errorf("stock disponible debería ser 5, got %d", p.GetStockDisponible())
	}
}

func TestLiberarStock(t *testing.T) {
	p := NewRopa("Camisa", 25.00, 10, "M", "Algodón")
	_ = p.Reservar(3)
	p.LiberarReserva(2)

	if p.GetStockDisponible() != 9 {
		t.Errorf("stock disponible debería ser 9 después de liberar reserva, got %d", p.GetStockDisponible())
	}
}
