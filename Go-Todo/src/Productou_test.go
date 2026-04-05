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

func TestIDArrancanDesdeUno(t *testing.T) {
	g := nuevoGen()
	p := NewLibroConGenerador("Test", 10.00, 1, "Autor", "Cat", g)
	if p.GetID() != 1 {
		t.Errorf("primer producto debería tener ID 1, got %d", p.GetID())
	}
}

func TestIDsCrecientes(t *testing.T) {
	g := nuevoGen()
	a := NewRopaConGenerador("A", 10.00, 1, "S", "Tela", g)
	b := NewRopaConGenerador("B", 10.00, 1, "M", "Tela", g)
	if b.GetID() <= a.GetID() {
		t.Errorf("ID de b (%d) debería ser mayor que a (%d)", b.GetID(), a.GetID())
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

func TestConfirmarComprar(t *testing.T) {
	p := NewElectronico("celu", 700.04, 5, "iphone", 12)
	_ = p.Reservar(2)
	_ = p.ConfirmarCompra(2)

	if p.GetStock() != 3 {
		t.Errorf("stock real debería ser 3 después de confirmar compra, got %d", p.GetStock())
	}

	if p.GetStockDisponible() != 3 {
		t.Errorf("stock disponible debería ser 3 después de confirmar compra, got %d", p.GetStockDisponible())
	}
}

func TestStockInsuficiente(t *testing.T) {
	p := NewElectronico("TV", 500.00, 1, "Samsung", 12)
	err := p.Reservar(5)
	if err == nil {
		t.Error("debería retornar error por stock insuficiente")
	}
}

func TestNoDeberiaEstarDisponible(t *testing.T) {
	p := NewLibro("Test", 10.00, 1, "Autor", "Cat")

	if !p.EstaDisponible() {
		t.Error("debería estar disponible antes de reservar")
	}
	_ = p.Reservar(1)
	if p.EstaDisponible() {
		t.Error("no debería estar disponible con todo el stock reservado")
	}
}
