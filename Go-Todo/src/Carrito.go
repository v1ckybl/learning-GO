package main

import (
	"fmt"
	"strings"
)

type Carrito struct {
	id        string
	items     map[uint64]ItemCompra // clave: ID único del producto
	descuento float64               // porcentaje 0-100
}

func NewCompra(id string) *Carrito {
	return &Carrito{
		id:    id,
		items: make(map[uint64]ItemCompra),
	}
}

// MÉTODO: AgregarProducto
// Go no tiene sobrecarga, pero podríamos agregar AgregarProductoConNota, etc.
func (c *Carrito) AgregarProducto(p IProducto, cantidad int) error {
	if !p.EstaDisponible() {
		return fmt.Errorf("producto '%s' sin stock", p.GetNombre())
	}
	err := p.Reservar(cantidad)
	if err != nil {
		return err
	}

	existente, existe := c.items[p.GetID()]
	if existe {
		existente.cantidad += cantidad
		existente.subtotal += p.GetPrecio() * float64(cantidad)
		c.items[p.GetID()] = existente
	} else {
		c.items[p.GetID()] = newItem(p, cantidad)
	}
	return nil
}

// MÉTODO: RemoverProducto — recibe el ID único del producto
func (c *Carrito) RemoverProducto(id uint64) {
	item, existe := c.items[id]
	if existe {
		item.producto.LiberarReserva(item.cantidad)
		delete(c.items, id)
	}

}

// MÉTODO: AplicarDescuento
func (c *Carrito) AplicarDescuento(porcentaje float64) {
	c.descuento = porcentaje
}

// MÉTODO: Total — calcula el total con descuento aplicado
func (c *Carrito) Total() float64 {
	subtotal := 0.0
	for _, item := range c.items {
		subtotal += item.subtotal
	}
	return subtotal * (1 - c.descuento/100)
}

// MÉTODO: CantidadItems
func (c *Carrito) CantidadItems() int {
	return len(c.items)
}

func (c *Carrito) ConfirmarCompra() error {
	for _, item := range c.items {
		err := item.producto.ConfirmarCompra(item.cantidad)
		if err != nil {
			return fmt.Errorf("error confirmando '%s': %w", item.producto.GetNombre(), err)
		}
	}
	c.items = make(map[uint64]ItemCompra) // vacía el carrito
	return nil
}

// MÉTODO: Resumen — genera texto descriptivo del carrito
func (c *Carrito) Resumen() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("🛒 Carrito [%s] — %d producto(s)\n", c.id, len(c.items)))
	sb.WriteString(strings.Repeat("─", 55) + "\n")
	for _, item := range c.items {
		sb.WriteString(fmt.Sprintf("  %-30s x%d  $%.2f\n",
			item.producto.GetNombre(), item.cantidad, item.subtotal))
	}
	sb.WriteString(strings.Repeat("─", 55) + "\n")
	sb.WriteString(fmt.Sprintf("  TOTAL: $%.2f\n", c.Total()))
	return sb.String()
}

// en vez de generar un resumen podriamos testear el total de la cantidad de items comprados
