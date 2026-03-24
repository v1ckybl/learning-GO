package main

type Carrito struct {
	Items []ItemCompra
}

func (c *Carrito) AddItem(item ItemCompra) { //modifica el carrito original con *
	c.Items = append(c.Items, item)
}

func (c *Carrito) Total() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.PrecioPorItem()
	}
	return total
}
