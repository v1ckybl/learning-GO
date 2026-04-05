package main

import (
	"fmt"
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

// Go no tiene sobrecarga, pero podríamos agregar AgregarProductoConNota
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

// RemoverProducto, recibe el ID único del producto
func (c *Carrito) RemoverProducto(id uint64) {
	item, existe := c.items[id]
	if existe {
		item.producto.LiberarReserva(item.cantidad)
		delete(c.items, id)
	}

}

func (c *Carrito) AplicarDescuento(porcentaje float64) {
	c.descuento = porcentaje
}

// calcula el total con descuento aplicado
func (c *Carrito) Total() float64 {
	subtotal := 0.0
	for _, item := range c.items {
		subtotal += item.subtotal
	}
	return subtotal * (1 - c.descuento/100)
}

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
