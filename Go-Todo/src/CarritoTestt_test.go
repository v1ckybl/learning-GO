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
		t.Errorf("stock real no deberia cambiar, got %d", laptop.GetStock())
	}
	if laptop.GetStockDisponible() != 3 {
		t.Errorf("stock disponible deberia ser 3, got %d", laptop.GetStockDisponible())
	}
}

func TestRemoverProductoLiberaReserva(t *testing.T) {
	carrito := NewCompra("T-002")
	libro := NewLibro("Dune", 35.00, 3, "Herbert", "Ficción")

	_ = carrito.AgregarProducto(libro, 2)
	carrito.RemoverProducto(libro.GetID())

	// al sacar del carrito, el stock vuelve a estar disponible
	if libro.GetStockDisponible() != 3 {
		t.Errorf("stock disponible deberia volver a 3, got %d", libro.GetStockDisponible())
	}
	if carrito.CantidadItems() != 0 {
		t.Errorf("carrito deberia estar vacío, got %d", carrito.CantidadItems())
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
		t.Error("carrito debería estar vacío despues de confirmar")
	}
	if remera.GetStock() != 3 {
		t.Errorf("stock real debería ser 3, got %d", remera.GetStock())
	}
}

func TestTotalSinDescuento(t *testing.T) {
	carrito := NewCompra("T-004")
	libro := NewLibro("SICP", 40.00, 10, "Abelson", "CS")

	_ = carrito.AgregarProducto(libro, 2)

	if carrito.Total() != 80.00 {
		t.Errorf("total incorrecto: got %.2f, want 80.00", carrito.Total())
	}
}

func TestTotalConDescuento(t *testing.T) {
	carrito := NewCompra("T-005")
	remera := NewRopa("Remera", 100.00, 10, "M", "Lino")

	_ = carrito.AgregarProducto(remera, 1)
	carrito.AplicarDescuento(20)

	if carrito.Total() != 80.00 {
		t.Errorf("total con descuento incorrecto: got %.2f, want 80.00", carrito.Total())
	}
}

func TestAgregarSinStockDisponible(t *testing.T) {
	carrito := NewCompra("T-006")
	p := NewElectronico("Auriculares", 60.00, 0, "Sony", 6)

	err := carrito.AgregarProducto(p, 1)
	if err == nil {
		t.Error("debería retornar error: producto sin stock")
	}
}

func TestAgregarMismoProductoDosVeces(t *testing.T) {
	carrito := NewCompra("T-008")
	libro := NewLibro("Harry Potter", 35.00, 5, "J.K. Rowling", "Fantasia")

	_ = carrito.AgregarProducto(libro, 1)
	_ = carrito.AgregarProducto(libro, 1)

	if carrito.CantidadItems() != 1 {
		t.Errorf("debería haber 1 item acumulado, got %d", carrito.CantidadItems())
	}
}

func TestConfirmarCompraConErrorRetorna(t *testing.T) {
	carrito := NewCompra("T-009")
	p := NewRopa("Short", 25.00, 1, "L", "Algodón")

	_ = carrito.AgregarProducto(p, 1)
	p.stock = 0 // forzamos inconsistencia

	err := carrito.ConfirmarCompra()
	if err == nil {
		t.Error("debería retornar error al confirmar con stock insuficiente")
	}
}

func TestInterfazProducto(t *testing.T) {
	g := nuevoGen()
	productos := []IProducto{
		NewElectronicoConGenerador("Monitor", 300.00, 2, "LG", 18, g),
		NewRopaConGenerador("Jean", 55.00, 7, "38", "Denim", g),
		NewLibroConGenerador("The Pragmatic Programmer", 42.00, 4, "Hunt & Thomas", "Dev", g),
	}

	carrito := NewCompra("T-007")
	for _, p := range productos {
		err := carrito.AgregarProducto(p, 1)
		if err != nil {
			t.Errorf("error inesperado agregando %s: %v", p.GetNombre(), err)
		}
	}

	if carrito.CantidadItems() != 3 {
		t.Errorf("esperaba 3 items, got %d", carrito.CantidadItems())
	}
}
