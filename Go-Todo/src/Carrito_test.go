package main

import (
	"testing"
)

func TestAgregarProductoReserva(t *testing.T) {
	carrito := NewCompra("T-001")
	laptop := NewElectronico("Laptop", 800.00, 5, "HP", 24)

	_ = carrito.AgregarProducto(laptop, 2)

	// stock real intacto, solo reservado
	if laptop.GetStock() != 5 {
		t.Errorf("stock real no debería cambiar, got %d", laptop.GetStock())
	}
	if laptop.GetStockDisponible() != 3 {
		t.Errorf("stock disponible debería ser 3, got %d", laptop.GetStockDisponible())
	}
}

func TestRemoverProductoLiberaReserva(t *testing.T) {
	carrito := NewCompra("T-002")
	libro := NewLibro("Dune", 35.00, 3, "Herbert", "Ficción")

	_ = carrito.AgregarProducto(libro, 2)
	carrito.RemoverProducto(libro.GetID())

	// al sacar del carrito, el stock vuelve a estar disponible
	if libro.GetStockDisponible() != 3 {
		t.Errorf("stock disponible debería volver a 3, got %d", libro.GetStockDisponible())
	}
	if carrito.CantidadItems() != 0 {
		t.Errorf("carrito debería estar vacío, got %d", carrito.CantidadItems())
	}
}

func TestConfirmarCompraYVaciaCarrito(t *testing.T) {
	carrito := NewCompra("T-003")
	remera := NewRopa("Remera", 100.00, 5, "M", "Lino")

	_ = carrito.AgregarProducto(remera, 2)
	err := carrito.ConfirmarCompra()

	if err != nil {
		t.Fatalf("no debería haber error al confirmar: %v", err)
	}
	if carrito.CantidadItems() != 0 {
		t.Error("carrito debería estar vacío después de confirmar")
	}
	if remera.GetStock() != 3 {
		t.Errorf("stock real debería ser 3, got %d", remera.GetStock())
	}
}
