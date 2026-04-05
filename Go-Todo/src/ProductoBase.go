package main

import (
	"fmt"
)

type ProductoBase struct { //clase abstracta que implementa la interfaz Producto, sirve como base para otros tipos de productos
	id             uint64
	nombre         string
	precio         float64
	stock          int
	stockReservado int
}

func nuevoProductoBase(nombre string, precio float64, stock int) ProductoBase {
	return nuevoProductoBaseConGenerador(nombre, precio, stock, idGeneraaado)
}

func nuevoProductoBaseConGenerador(nombre string, precio float64, stock int, g *GenerarID) ProductoBase {
	return ProductoBase{
		id:     g.generarID(),
		nombre: nombre,
		precio: precio,
		stock:  stock,
	}
}

func (p *ProductoBase) GetID() uint64 {
	return p.id
}

func (p *ProductoBase) GetNombre() string {
	return p.nombre
}

func (p *ProductoBase) GetPrecio() float64 {
	return p.precio
}

func (p *ProductoBase) GetStock() int {
	return p.stock
}

func (p *ProductoBase) GetStockDisponible() int {
	return p.stock - p.stockReservado
}

// EstaDisponible considera las reservas activas
func (p *ProductoBase) EstaDisponible() bool {
	return p.GetStockDisponible() > 0
}

// reserva stock para una compra pendiente, no afecta el stock real hasta confirmar la compra
func (p *ProductoBase) Reservar(cantidad int) error {
	if p.GetStockDisponible() < cantidad {
		return fmt.Errorf("stock insuficiente: disponible %d, pedido %d",
			p.GetStockDisponible(), cantidad)
	}
	p.stockReservado += cantidad
	return nil
}

// libera una reserva sin afectar el stock real, útil para cancelar compras pendientes
func (p *ProductoBase) LiberarReserva(cantidad int) {
	p.stockReservado -= cantidad
}

// agora si descuenta el stock real xq se confirma la compraa
func (p *ProductoBase) ConfirmarCompra(cantidad int) error {
	if p.stock < cantidad {
		return fmt.Errorf("stock real insuficiente: %d", p.stock)
	}
	p.stock -= cantidad
	p.stockReservado -= cantidad
	return nil
}
